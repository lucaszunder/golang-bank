package account

import c "github.com/lucaszunder/banco/customer"

type Account struct {
	Owner                       c.Owner
	AgencyNumber, AccountNumber int
	balance                     float64
}

func (c *Account) Withdraw(withdrawValue float64) string {
	canWithdraw := withdrawValue <= c.balance && withdrawValue > 0

	if canWithdraw {
		c.balance -= withdrawValue
		return "Saque realizado com sucesso"
	} else {
		return "Saldo insuficiente"
	}
}

func (c *Account) Deposit(depositValue float64) (string, float64) {
	canDeposit := depositValue > 0

	if canDeposit {
		c.balance += depositValue
		return "Foi realizado um depósito com sucesso", c.balance
	} else {
		return "Não foi possível realizar um depósito nesta conta", c.balance
	}
}

func (c *Account) Transfer(transferValue float64, destinyAccount *Account) bool {
	canTransfer := transferValue <= c.balance && transferValue > 0

	if canTransfer {
		c.balance -= transferValue
		destinyAccount.Deposit(transferValue)
		return true
	} else {
		return false
	}
}

func (c *Account) Balance() float64 {
	return c.balance
}

// declaracões de struct
// var account2 *Account
// account2 = new(Account)
// account2.owneO = "Cris"B
