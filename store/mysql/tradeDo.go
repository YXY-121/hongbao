package mysql

type Trade struct {
	ID           int     `gorm:"column:good_id" json:"id"`
	TradeId      int     `gorm:"column:trade_id" json:"trade_id"`
	AccountId    int     `gorm:"column:account_id" json:"account_id"`
	SenderId     int     `gorm:"column:sender_id" json:"sender_id"`
	SenderName   string  `gorm:"column:sender_name" json:"sender_name"`
	AccepterId   int     `gorm:"column:accepter_id" json:"accepter_id"`
	AccepterName string  `gorm:"column:accepter_name" json:"accepter_name"`
	Amount       float64 `gorm:"column:amount" json:"amount"`
	Balance      float64 `gorm:"column:balance" json:"balance"`
}
