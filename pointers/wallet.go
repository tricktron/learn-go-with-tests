package pointers

import (
	"errors"
	"fmt"
)

var errInsufficientFunds = errors.New("insufficient funds")

type Bitcoin int

type Wallet struct {
	balance Bitcoin
}

func NewWallet(initBalance Bitcoin) *Wallet {
	return &Wallet{balance: initBalance}
}

func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return errInsufficientFunds
	}

	w.balance -= amount

	return nil
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}
