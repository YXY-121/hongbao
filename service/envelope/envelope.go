package envelope

import (
	"context"
	"github.com/shopspring/decimal"
	apis "hongbao/apis/api/envelope"
	"hongbao/store"
	"hongbao/store/mysql"
	"math/rand"
)

type envelopeService struct {
	envelopeDao store.OperateEnvelope
}
type EnvelopeDTO struct {
	id int

}
//总红包
type EnvelopeVo struct {
	EnvelopeId int//红包id
	SenderId int //发送者
	blessing string //祝福
	TotalAmount decimal.Decimal //总金额
}
//红包项
type EnvelopeItemVo struct {
	EnvelopeId int//所属大红包id
	EnvelopeItemId int//红包item
	AccepterId int//接收者
	ItemAmount decimal.Decimal//单个红包金额
}

//分配红包红包类型生成响应的红包数量、以及money\生成红包
func (e *envelopeService) DivideHongbao(ctx context.Context, envelope *apis.Envelope) {
	////过期时间
	//envelopeId := uuid.UUID{}
	//
	//expiredTime := envelope.CreateTime.Add(24 * time.Hour)
	//updateTime := envelope.CreateTime
	//if envelope.EnvelopeType == 1 { //普通红包用一个红包id搞定
	//	//单个金额
	//	singleAmount := envelope.TotalAmount / float64(envelope.Quantity)
	//} else { //随机红包
	//	randomCreateHongbao(envelope.TotalAmount, envelope.Quantity)
	//}

}

//提前生成好红包
func RandomCreateHongbao(totalAmount float64, quantity int) []decimal.Decimal{
	now := decimal.NewFromFloat(totalAmount)
	hongbaoList := make([]decimal.Decimal, quantity)

	for i := 0; i < quantity-1; i++ {
		max := decimal.NewFromFloat((totalAmount / float64(quantity)) * 2)
		ran := decimal.NewFromFloat(rand.Float64())
		value := max.Mul(ran)
		if value.LessThan(decimal.NewFromFloat(0.01)) {
			value = decimal.NewFromFloat(0.01)

		} else {
			value = value.Mul(decimal.NewFromInt(100)).Div(decimal.NewFromInt(100))

		}
		hongbaoList = append(hongbaoList, value)
		quantity--
		now.Sub(value)

	}

	hongbaoList = append(hongbaoList, now.Mul(decimal.NewFromInt(100)).Div(decimal.NewFromInt(100)))
	return hongbaoList
}

//红包的首页：红包的金额、红包数量、发红包的人名字、红包的主福、当前剩余总书、:抢到的人的名字、金额
func (e *envelopeService) GetHongbaoMessage(ctx context.Context, envelopeId int) (*mysql.EnvelopeDo, error) {
	//获取大红包

	//获取红包小项们

	//查看缓存是否存
	//RedisEnvelope,err:=mysql.RedisDB.Do("get",envelopeId)

	//if err!=nil {
	//	////获取红包小项
	//	//RedisEnvelopeItem,err:=mysql.RedisDB.Do("get",envelopeId)
	//	//
	//	//return RedisEnvelope;
	//}

	//存在就在数据库取

	//不存在就去数据库取
	envelope, err := e.envelopeDao.SelectEnvelopeById(ctx, envelopeId)
	return envelope, err
}


