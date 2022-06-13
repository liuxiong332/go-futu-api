package go_futu_api

import (
	"time"

	"github.com/liuxiong332/go-futu-api/pb/initconnect"
	"github.com/liuxiong332/go-futu-api/pb/keepalive"
)

func (me *Client) InitConnect() (resp *initconnect.Response, err error) {
	req := &initconnect.Request{
		C2S: &initconnect.C2S{
			ClientID:  &ClientID,
			ClientVer: &ClientVer,
		},
	}
	resp = &initconnect.Response{}
	err = me.DoRequest(uint32(1001), req, resp)
	return
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
