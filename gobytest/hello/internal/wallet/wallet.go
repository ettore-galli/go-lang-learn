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

func SearchOnMap(searchMap map[string]string, key string) string {
	return searchMap[key]
}

type Dictionary map[string]string

func (dict Dictionary) Search(key string) (string, error) {
	value, exists := dict[key]
	if !exists {
		return "", fmt.Errorf("non trovato %s", key)
	}
	return value, nil
}

func (dict Dictionary) AddEntry(key string, value string) error {
	_, exists := dict[key]
	if exists {
		return fmt.Errorf("entry [%s] exists", key)
	}
	dict[key] = value

	return nil
}

func (dict Dictionary) UpdateEntry(key string, value string) error {
	_, exists := dict[key]
	if !exists {
		return fmt.Errorf("entry [%s] does not exist", key)
	}
	dict[key] = value

	return nil
}

func (dict Dictionary) DeleteEntry(key string) error {
	delete(dict, key)
	return nil
}
