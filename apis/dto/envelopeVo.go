package vo

import (
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"time"
)

type EnvelopeVo struct {
	EnvelopeId  uuid.UUID       //红包id
	SenderId    string          //发送者
	Blessing    string          //祝福
	Number      int             //个数
	TotalAmount decimal.Decimal //总金额
	Status      string             //类型
	CreateTime  time.Time       //现在的时间

}

func (e EnvelopeVo) Deadline() (deadline time.Time, ok bool) {
	panic("implement me")
}

func (e EnvelopeVo) Done() <-chan struct{} {
	panic("implement me")
}

func (e EnvelopeVo) Err() error {
	panic("implement me")
}

func (e EnvelopeVo) Value(key interface{}) interface{} {
	panic("implement me")
}
