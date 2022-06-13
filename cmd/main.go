package main

import (
	"fmt"
	"log"
	"time"

	gf "github.com/liuxiong332/go-futu-api"
	"github.com/liuxiong332/go-futu-api/pb/qotcommon"
)

func main() {
	// 初始化opend的配置
	cli := gf.NewClient("0.0.0.0:11111")

	// 启动
	err := cli.Run()
	if err != nil {
		panic(err)
	}

	// 建立tcp连接
	resp, err := cli.InitConnect()
	if err != nil {
		fmt.Println(err)
	}

	// 循环发送keepalive请求
	keepAliveInterval := resp.S2C.KeepAliveInterval
	go func() {
		for {
			d := time.Duration(int64(*keepAliveInterval)) * time.Second
			time.Sleep(d)
			cli.KeepAlive()
		}
	}()

	// 获取美股所有板块信息
	respGetPlateSet, err := cli.GetPlateSet(
		qotcommon.QotMarket_QotMarket_US_Security,
		qotcommon.PlateSetType_PlateSetType_All,
	)

	if err != nil {
		panic(err)
	}

	for idx, item := range respGetPlateSet.S2C.PlateInfoList {
		log.Println("plate:", idx, *item.Plate.Market, *item.Plate.Code, *item.Name)
	}

	// 获取nasdaq所有股票列表
	market := int32(11)
	code := "NASDAQ"
	respGetPlateSecurity, err := cli.GetPlateSecurity(&qotcommon.Security{
		Market: &market,
		Code:   &code,
	})
	if err != nil {
		panic(err)
	}

	for idx, item := range respGetPlateSecurity.S2C.StaticInfoList {
		log.Println("security:", idx, *item.Basic.Name, *item.Basic.Security.Code)
	}

}
