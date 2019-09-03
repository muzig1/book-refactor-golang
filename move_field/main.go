package main

import "fmt"

/*
手法：Move Field - 移动字段

注意：
1. 重构字段位置的时候，注意数据库对应也应该发生变化
*/

func main() {
	acc := new(Account)
	acc.Type = AccountType{
		Type:         "vip",
		InterestRate: 0.8,
	}
	money := acc.InterestForAmount(1*1e5, 100)
	fmt.Println(money)
}

type (
	Account struct {
		Type AccountType
	}

	AccountType struct {
		Type         string
		InterestRate float32
	}
)

func (at *Account) InterestForAmount(amount float32, days int) float32 {
	return at.Type.interestForAmount(amount, days)
}

func (at *AccountType) interestForAmount(amount float32, days int) float32 {
	return at.InterestRate * amount * float32(days) / 365
}
