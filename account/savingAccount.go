package account

import (
	c "github.com/lucaszunder/banco/customer"
)

type SavingAccount struct {
	Owner                                  c.Owner
	AgencyNumber, AccountNumber, Operation int
	balance                                float64
}

func (c *SavingAccount) Withdraw(withdrawValue float64) string {
	canWithdraw := withdrawValue <= c.balance && withdrawValue > 0

	if canWithdraw {
		c.balance -= withdrawValue
		return "Saque realizado com sucesso"
	} else {
		return "Saldo insuficiente"
	}
}

func (c *SavingAccount) Deposit(depositValue float64) (string, float64) {
	canDeposit := depositValue > 0

	if canDeposit {
		c.balance += depositValue
		return "Foi realizado um depósito com sucesso", c.balance
	} else {
		return "Não foi possível realizar um depósito nesta conta", c.balance
	}
}

func (c *SavingAccount) Balance() float64 {
	return c.balance
}
