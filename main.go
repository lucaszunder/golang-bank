package main

import (
	"fmt"

	a "github.com/lucaszunder/banco/account"
	c "github.com/lucaszunder/banco/customer"
)

func main() {
	owner := c.Owner{Name: "Lucas Zunder", CPF: "000888123", Job: "Engenheiro de software"}
	account := a.Account{Owner: owner, AgencyNumber: 589, AccountNumber: 4397}
	account.Deposit(100)

	fmt.Println(account.Balance())
}
