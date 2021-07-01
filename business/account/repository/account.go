package repository

import (
	"context"
	"gorm.io/gorm"
	"hongbao/business/account/repository/mysql"
)

type OperateAccount interface {
		//定义数据库的接口
	SelectAccountById(ctx context.Context,Id int64)(,error)
}

type Account struct {
	db *gorm.DB
}

func NewAccount(db *gorm.DB) *Account{
	return &Account{
		db: db,
	}
}

func (account *Account)SelectAccountById(ctx context.Context,Id int64) error {
	accountDo:=&mysql.AccountDo{}
	err:=account.db.Where("id=").Row()
	if err!=nil {
		return nil
	}

	
}


