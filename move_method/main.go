package main

import "fmt"

/*
手法：Move Method - 迁移函数

目的：1. 减少一个类的职责 2. 职责划分的更加的明确
*/

func main() {
	acc := new(Account)
	acc.DaysOverdrawn = 10
	acc.Type = AccountTypePremium

	charge := acc.BankCharge()
	fmt.Println("charge:%v", charge)
}

const (
	AccountTypePremium AccountType = "premium"
)

type (
	Account struct {
		Type          AccountType
		DaysOverdrawn int
	}

	AccountType string
)

func (a *Account) overdraftCharge() float32 {
	if a.Type == AccountTypePremium {
		result := float32(10)
		if a.DaysOverdrawn > 7 {
			result += (float32(a.DaysOverdrawn) - 7) * 0.85
		}
		return result
	} else {
		return float32(a.DaysOverdrawn) * 1.75
	}
}

func (a *Account) BankCharge() float32 {
	result := float32(4.5)
	if a.DaysOverdrawn > 0 {
		result += a.overdraftCharge()
	}
	return result
}
