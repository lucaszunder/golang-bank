package main

import (
	"fmt"

	a "github.com/lucaszunder/banco/account"
	c "github.com/lucaszunder/banco/customer"
)

func payBills(account verifyAccountType, value float64) {
	account.Withdraw(value)
}

type verifyAccountType interface {
	Withdraw(value float64) string
}

func main() {
	owner := c.Owner{Name: "Lucas Zunder", CPF: "000888123", Job: "Engenheiro de software"}
	account := a.Account{Owner: owner, AgencyNumber: 589, AccountNumber: 4397}
	account.Deposit(100)

	payBills(&account, 50.0)

	fmt.Println(account.Balance())
}
