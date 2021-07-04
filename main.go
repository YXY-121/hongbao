package main

import (
	"fmt"
	"github.com/shopspring/decimal"
	"hongbao/store/mysql"
	"math/rand"
)

func main()  {
		mysql.GetDB()
		mysql.GetRedis()

}
func RandomCreateHongbao(totalAmount float64, quantity int) []decimal.Decimal{
	now := decimal.NewFromFloat(totalAmount)
	hongbaoList := make([]decimal.Decimal, quantity)
	len:= quantity
	for i := 0; i < len-1; i++ {
		max := decimal.NewFromFloat((totalAmount / float64(quantity)) * 2)
		ran := decimal.NewFromFloat(rand.Float64())
		value := max.Mul(ran)
		fmt.Println(value)
		if value.LessThan(decimal.NewFromFloat(0.01)) {
			value = decimal.NewFromFloat(0.01)

		} else {
			value = value.Mul(decimal.NewFromInt(100)).Div(decimal.NewFromInt(100))

		}
		hongbaoList = append(hongbaoList, value)
		quantity--
		now=now.Sub(value)

	}

	hongbaoList = append(hongbaoList, now.Mul(decimal.NewFromInt(100)).Div(decimal.NewFromInt(100)))
	return hongbaoList
}
