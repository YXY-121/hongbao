package apis

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	vo "hongbao/apis/dto"
	"hongbao/service/envelope"
	"strconv"
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

func CreateEnvelope(c *gin.Context)  {

	number, _ := strconv.Atoi(c.PostForm("number"))
	totalAmount, _ := decimal.NewFromString(c.PostForm("totalAmount"))
	e:=vo.EnvelopeVo{
		EnvelopeId:  uuid.New(),
		Blessing:    c.PostForm("blessing"),
		SenderId:    c.PostForm("senderId"),
		Status:      c.PostForm("status"),
		Number:      number,
		TotalAmount: totalAmount,
		CreateTime:  time.Now(),
	}

	eService := envelope.EnvelopeService{}
	eService.DivideHongbao(&e)
}
func Hi(c *gin.Context)  {

	fmt.Println(c.QueryMap("a"))
	//c.Query("a")
}



