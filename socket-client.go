package go_futu_api

import (
	"net"

	"github.com/liuxiong332/go-futu-api/pb/initconnect"
	"google.golang.org/protobuf/proto"
)

type SocketClient struct {
	addr string
	conn net.Conn
}

func NewSocketClient(addr string) *SocketClient {
	return &SocketClient{
		addr: addr,
	}
}

func (me *SocketClient) Connect() error {
	conn, err := net.Dial("tcp", me.addr)
	if err != nil {
		panic(err)
	}
	me.conn = conn
	return nil
}

func (me *SocketClient) HandleWrite() {

	pack := &FutuPack{}
	pack.SetProtoID(uint32(1001))
	// pack.SetBody([]byte("现在时间是:" + time.Now().Format("2006-01-02 15:04:05")))
	// pack.SetBody([]byte("test msg from client"))

	ClientVer := int32(307)
	ClientID := "test123456"

	fut := &initconnect.Request{
		C2S: &initconnect.C2S{
			ClientID:  &ClientID,
			ClientVer: &ClientVer,
		},
	}
	// fut.ClientVer = &ClientVer
	// fut.ClientID = &ClientID
	body, err := proto.Marshal(fut)
	if err != nil {
		//log.Fatal("marshaling error: ", err)
	}
	pack.SetBody(body)

	pack.Send(me.conn)

}

//func (me *SocketClient) handleMessage() {}

//func (me *Client)
