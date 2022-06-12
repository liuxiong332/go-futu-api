package go_futu_api

import (
	"github.com/futuopen/ftapi4go/pb/qotcommon"
	"github.com/futuopen/ftapi4go/pb/qotgetbasicqot"
	"github.com/futuopen/ftapi4go/pb/qotgetplatesecurity"
	"github.com/futuopen/ftapi4go/pb/qotgetplateset"
	"github.com/futuopen/ftapi4go/pb/qotgetsecuritysnapshot"
	"github.com/futuopen/ftapi4go/pb/qotgetstaticinfo"
	"github.com/futuopen/ftapi4go/pb/qotgettradedate"
)

func (me *Client) GetBasicQot(list []*qotcommon.Security) (resp *qotgetbasicqot.Response, err error) {
	req := &qotgetbasicqot.Request{
		C2S: &qotgetbasicqot.C2S{
			SecurityList: list,
		},
	}
	resp = &qotgetbasicqot.Response{}
	err = me.DoRequest(uint32(3004), req, resp)
	return
}

func (me *Client) GetTradeDate(
	market qotcommon.QotMarket,
	beginTime string,
	endTime string,
) (resp *qotgettradedate.Response, err error) {
	marketInt32 := int32(market)
	req := &qotgettradedate.Request{
		C2S: &qotgettradedate.C2S{
			Market:    &marketInt32,
			BeginTime: &beginTime,
			EndTime:   &endTime,
		},
	}
	resp = &qotgettradedate.Response{}
	err = me.DoRequest(uint32(3200), req, resp)
	return
}

func (me *Client) GetSecuritySnapshot(list []*qotcommon.Security) (resp *qotgetsecuritysnapshot.Response, err error) {
	req := &qotgetsecuritysnapshot.Request{
		C2S: &qotgetsecuritysnapshot.C2S{
			SecurityList: list,
		},
	}
	resp = &qotgetsecuritysnapshot.Response{}
	err = me.DoRequest(uint32(3203), req, resp)
	return
}

func (me *Client) GetPlateSet(
	market qotcommon.QotMarket,
	plateSetType qotcommon.PlateSetType,
) (resp *qotgetplateset.Response, err error) {
	marketParam := int32(market)
	plateSetTypeParam := int32(plateSetType)
	req := &qotgetplateset.Request{
		C2S: &qotgetplateset.C2S{
			Market:       &marketParam,
			PlateSetType: &plateSetTypeParam,
		},
	}
	resp = &qotgetplateset.Response{}
	err = me.DoRequest(uint32(3204), req, resp)
	return
}

func (me *Client) GetPlateSecurity(plate *qotcommon.Security) (resp *qotgetplatesecurity.Response, err error) {
	req := &qotgetplatesecurity.Request{
		C2S: &qotgetplatesecurity.C2S{
			Plate: plate,
		},
	}
	resp = &qotgetplatesecurity.Response{}
	err = me.DoRequest(uint32(3205), req, resp)
	return
}

func (me *Client) GetStaticInfo(
	market *qotcommon.QotMarket,
	secType *qotcommon.SecurityType,
	list []*qotcommon.Security,
) (resp *qotgetstaticinfo.Response, err error) {
	c2s := &qotgetstaticinfo.C2S{}
	if market != nil {
		marketInt32 := int32(*market)
		c2s.Market = &marketInt32
	}
	if secType != nil {
		secTypeInt32 := int32(*secType)
		c2s.SecType = &secTypeInt32
	}
	if len(list) != 0 {
		c2s.SecurityList = list
	}
	req := &qotgetstaticinfo.Request{
		C2S: c2s,
	}
	resp = &qotgetstaticinfo.Response{}
	err = me.DoRequest(uint32(3202), req, resp)
	return
}
