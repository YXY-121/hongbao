package apis

import (
	"context"
	"github.com/gin-gonic/gin"
	"hongbao/service"
	"time"
)

type Envelope struct {
	EnvelopeType int //红包类型
	UserId int
	TotalAmount float64 //总金额
//	RemainedAmount float64//总的剩金额
//	OneAmount float64 //每个金额
	Quantity int //红包数量
//	ExpiredTime time.Time//过期时间
	CreateTime time.Time
//	UpdateTime time.Time
}

func CreateEnvelope(c *gin.Context,ctx context.Context,envelope *Envelope)  {
	//e:=c.BindJSON(&Envelope{})
	service.EnvelopeService.DivideHongbao(ctx,envelope)



}



