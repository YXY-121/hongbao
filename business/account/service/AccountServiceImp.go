package service

import "hongbao/business/account/repository"

type AccountService struct {
	accountDo repository.OperateAccount
}

func (a *AccountService)GetAccountById(id int)  {

}
func (a *AccountService)GetAccountBalanceById(id int) float64{
	
}
func (a *AccountService)IsLitterThenOperatingMone(id int,float642 float64) bool  {
	
}
func (a *AccountService)Transfer(nowId int,targetId int,money float64)  {

}

func (a *AccountService)AddAccountBalance(id int,money float64)  {

}

func (a *AccountService)DecreaseAccountBalance(id int,money float64)  {

}



