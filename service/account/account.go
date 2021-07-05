package account

import (
	"context"
	"hongbao/store"
)

type AccountService struct {
	account store.OperateAccount
}

func (accountService *AccountService) GetAccountById(ctx context.Context, id int) {
	//判断id是否合法、存在

	//判断是否在缓存中存在

	//不在就去数据库获取
	//accountService.account.SelectAccountById()
}

//查询用户Id对应账户里的余额
func (account *AccountService) GetAccountBalanceById(ctx context.Context,id int) float64 {
	//判断缓存的账户 是否存在

	//缓存存在就返回账户中的余额

	//不存在就调用GetAccountById
	return 1.1
}

//判断余额是否小于当前操作金e
func (account *AccountService) IsLitterThenOperatingMone(ctx context.Context,id int, float642 float64) bool {
	//从缓存中获取

	//如果缓存中没有就去拿
	return true
}

//转账
func (account *AccountService) Transfer(ctx context.Context,nowId int, targetId int, money float64) {
	//转入中间商 //中间商再转出来
}

//账户加钱
func (account *AccountService) AddAccountBalance(ctx context.Context,id int, money float64) {

}

//账户扣款
func (account *AccountService) DecreaseAccountBalance(ctx context.Context,id int, money float64) {

}
