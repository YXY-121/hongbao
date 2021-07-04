package mysql

import (
	"github.com/google/uuid"
	"time"
)

type AccountDo struct {
	Id             uint64     `gorm:"column:id;primary_key" json:"id"`
	AccountId      uuid.UUID  `gorm:"column:account_id" json:"account_id"`
	AccountName    string     `gorm:"column:account_name" json:"account_name"`
	AccountType    string     `gorm:"column:account_type" json:"account_type"`
	AccountBalance int        `gorm:"column:account_balance" json:"account_balance"`
	UserId         uint64     `gorm:"column:user_id" json:"user_id"`
	UserName       string     `gorm:"column:user_name" json:"user_name"`
	Status         string     `gorm:"column:status" json:"status"`
	CreateTime     *time.Time `gorm:"column:create_time" json:"create_time"`
	UpdateTime     *time.Time `gorm:"column:update_time" json:"update_time"`
}

func (a AccountDo) TableName() string {
	return "account"
}

//respository里写服务
