package service

import (
	"context"
	apis "hongbao/apis/api/envelope"
)

type EnvelopeService interface {




	//分配红包红包类型生成响应的红包数量、以及money\生成红包
	DivideHongbao(envelope *apis.Envelope)

	//红包的首页：红包的金额、红包数量、发红包的人名字、红包的主福、当前剩余总书、:抢到的人的名字、金额
	GetHongbaoMessage(ctx context.Context,envelopeId int)

}