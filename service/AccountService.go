package service

//账号服务
type AccountService interface {
	//根据用户id获取其对应的账号
	GetAccountById(id int)

	//查询用户Id对应账户里的余额
	GetAccountBalanceById(id int) float64

	//判断余额是否小于当前操作金e
	IsLitterThenOperatingMone(id int, float642 float64) bool

	//转账
	Transfer(nowId int, targetId int, money float64)

	//账户加钱
	AddAccountBalance(id int, money float64)

	//账户扣款
	DecreaseAccountBalance(id int, money float64)
}
