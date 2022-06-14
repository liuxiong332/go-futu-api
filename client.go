package go_futu_api

import (
	"bufio"
	"bytes"
	"crypto/rsa"
	"encoding/binary"
	"fmt"
	"log"
	"sync"

	"google.golang.org/protobuf/proto"
)

var (
	ClientVer = int32(100)
	ClientID  = "go-futu-api-100"
)

type Client struct {
	packChanMap   sync.Map
	socket        *SocketClient
	rsaPrivateKey *rsa.PrivateKey
	aesCipher     *AESCipher
}

func NewClient(addr string) *Client {
	return &Client{
		socket: NewSocketClient(addr),
	}
}

func (me *Client) SetRSAPrivateKey(key string) error {
	rsaPrivateKey, err := GenRsaPair([]byte(key))
	if err != nil {
		return err
	}
	me.rsaPrivateKey = rsaPrivateKey
	return nil
}

func (me *Client) Run() (err error) {
	err = me.socket.Connect()
	if err != nil {
		return
	}
	go func() {
		log.Println("futu api client run, addr:", me.socket.addr)
		me.WatchMessage()
	}()
	return
}

func (me *Client) SyncDo(requestPack *FutuPack) (*FutuPack, error) {
	ch := make(chan *FutuPack)
	defer close(ch)
	if err := requestPack.Send(me.socket.conn); err != nil {
		return nil, err
	}
	sid := requestPack.GetSerialNoStr()
	me.packChanMap.Store(sid, ch)
	defer me.packChanMap.Delete(sid)

	//log.Println("send msg", sid)
	responsePack := <-ch
	return responsePack, nil
}

func (me *Client) DoRequest(protoId uint32, request proto.Message, response proto.Message) error {
	pack := &FutuPack{}
	pack.SetProtoID(protoId)
	body, err := proto.Marshal(request)
	if err != nil {
		return err
	}

	if protoId == 1001 {
		// 初始化连接，使用rsa进行加密
		if me.rsaPrivateKey != nil {
			body, err = RsaEncrypt(me.rsaPrivateKey, body)
			if err != nil {
				return err
			}
		}
	} else if me.aesCipher != nil {
		// 其他请求使用AES进行加密
		body = me.aesCipher.Encrypt(body)
	}
	pack.SetBody(body)

	respPack, err := me.SyncDo(pack)
	if err != nil {
		return err
	}
	b := respPack.GetBody()
	return proto.Unmarshal(b, response)
}

func (me *Client) WatchMessage() {
	scanner := bufio.NewScanner(me.socket.conn)
	scanner.Split(func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		if !atEOF && data[0] == 'F' {
			if len(data) > 44 {
				length := uint32(0)
				binary.Read(bytes.NewReader(data[12:16]), binary.LittleEndian, &length)
				if int(length)+4 <= len(data) {
					return int(length) + 44, data[:int(length)+44], nil
				}
			}
		}
		return
	})
	//buf := make([]byte, 64 * 1024)
	buf := make([]byte, 0, 64*1024)
	scanner.Buffer(buf, 1024*1024)
	//scanner.Buffer(buf, bufio.MaxScanTokenSize)

	for scanner.Scan() {
		scannedPack := new(FutuPack)
		if err := scannedPack.Unpack(bytes.NewReader(scanner.Bytes())); err != nil {
			return
		}

		// 初始化连接，使用rsa进行解密
		if scannedPack.nProtoID == 1001 {
			if me.rsaPrivateKey != nil {
				body, err := RsaDecrypt(me.rsaPrivateKey, scannedPack.body)
				if err != nil {
					fmt.Printf("Invalid init content reply error: %v\n", err)
				}
				scannedPack.body = body
			}
		} else if me.aesCipher != nil {
			// 其他请求使用AES进行解密
			scannedPack.body = me.aesCipher.Decrypt(scannedPack.body)
		}

		sid := scannedPack.GetSerialNoStr()
		chInterface, ok := me.packChanMap.Load(sid)
		if ok {
			ch := chInterface.(chan *FutuPack)
			ch <- scannedPack
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("无效数据包", err)
	}
}
