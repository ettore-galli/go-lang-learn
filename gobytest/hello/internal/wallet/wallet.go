package wallet

import (
	"fmt"
)

type Bitcoin int

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

type Wallet struct {
	balance Bitcoin
}

func (wallet *Wallet) Deposit(amount Bitcoin) {
	(*wallet).balance += amount
}

func (wallet *Wallet) Withdraw(amount Bitcoin) error {
	if amount > (*wallet).Balance() {
		return fmt.Errorf("cannot withdraw %s from a balance of %s", amount, (*wallet).Balance())
	}
	(*wallet).balance -= amount
	return nil
}

func (wallet Wallet) Balance() Bitcoin {
	return wallet.balance
}
