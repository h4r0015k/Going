package wallet

import (
	"errors"
	"fmt"
)

type Bitcoin int

type Wallet struct {
	balance Bitcoin
}

var ErrInsufficientFunds = "insufficient balance"

func (w *Wallet) Deposit(amount Bitcoin) {
	fmt.Printf("address of wallet in Deposit is %p\n", w)
	w.balance += amount
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return errors.New(ErrInsufficientFunds)
	}

	w.balance -= amount
	return nil
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}
