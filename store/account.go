package store

import (
	"context"
	"hongbao/store/mysql"
)

type OperateAccount interface {
	//定义数据库的接口
	SelectAccountById(ctx context.Context, Id int64) (*mysql.AccountDo, error)
}


//func (account *OperateAccount) SelectAccountById(ctx context.Context, Id int64)  {
//	//
//	//err := db.Where("id=").Row()
//	//if err != nil {
//	//	return nil
//	//}
//	//return err
//}
