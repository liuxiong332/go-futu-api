package go_futu_api

import (
	"time"

	"github.com/liuxiong332/go-futu-api/pb/common"
	"github.com/liuxiong332/go-futu-api/pb/initconnect"
	"github.com/liuxiong332/go-futu-api/pb/keepalive"
)

func (me *Client) InitConnect() (*initconnect.Response, error) {
	packetEncAlgo := int32(common.PacketEncAlgo_PacketEncAlgo_None)
	if me.rsaPrivateKey != nil {
		packetEncAlgo = int32(common.PacketEncAlgo_PacketEncAlgo_AES_CBC)
	}

	req := &initconnect.Request{
		C2S: &initconnect.C2S{
			ClientID:      &ClientID,
			ClientVer:     &ClientVer,
			PacketEncAlgo: &packetEncAlgo,
		},
	}
	resp := &initconnect.Response{}
	if err := me.DoRequest(uint32(1001), req, resp); err != nil {
		return nil, err
	}
	if *resp.RetType == int32(common.RetType_RetType_Succeed) && me.rsaPrivateKey != nil && resp.S2C.ConnAESKey != nil {
		aesCipher, err := NewAESCipher([]byte(*resp.S2C.ConnAESKey), []byte(*resp.S2C.AesCBCiv))
		if err != nil {
			return resp, err
		}
		me.aesCipher = aesCipher
	}
	return resp, nil
}

func (me *Client) KeepAlive() (resp *keepalive.Response, err error) {
	t := int64(time.Now().Second())
	req := &keepalive.Request{
		C2S: &keepalive.C2S{
			Time: &t,
		},
	}
	resp = &keepalive.Response{}
	err = me.DoRequest(uint32(1004), req, resp)
	return
}
