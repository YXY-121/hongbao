package mysql

import (
	"context"
	"time"
)

type EnvelopeDo struct {
	ID           int    `gorm:"column:id" json:"id"`                        //自增id
	EnvelopeId   int    `gorm:"column:envelope_id" json:"envelope_id"`      //红包id
	EnvelopeType string `gorm:"column:envelope_type" json:"envelope_type" ` //	红包类型，1是普通红包 2是随机红包

	UserId         int        `gorm:"column:user_id" json:"user_id"`                 //	发红包者id
	UserName       string     `gorm:"column:user_name" json:"user_name"`             //发红包者名字
	TotalAmount    float64    `gorm:"column:total_amount" json:"total_amount"`       //总金额
	OneAmount      float64    `gorm:"column:one_amount" json:"one_amount"`           //单个红包金额
	Number         int        `gorm:"column:number" json:"number"`                   //红包个数
	RemainedAmount float64    `gorm:"column:remained_amount" json:"remained_amount"` //剩余金额
	Status         int        `gorm:"column:status" json:"status"`                   //0可用 1过期不可用
	ExpiredTime    time.Time `gorm:"column:expired_time" json:"expired_time"`       //过期时间
	CreateTime     time.Time `gorm:"column:create_time" json:"create_time"`         //	创建时间
	UpdateTime     time.Time `gorm:"column:update_time" json:"update_time"`         //修改时间

	//

}

func (a EnvelopeDo) TableName() string {
	return "envelope"
}

func (a EnvelopeDo) SelectEnvelopeById(ctx context.Context, Id int) (*EnvelopeDo, error) {
	hongbao := EnvelopeDo{}
	err := DB.Where("envelope_id=?", Id).First(&hongbao).Error
	if err != nil {
		return nil, nil
	}
	return &hongbao, nil

}
