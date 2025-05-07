package wallet

import "fmt"

type Wallet struct {
	balance int
}

func (w *Wallet) Deposit(amount int) {
	fmt.Printf("address of wallet in Deposit is %p\n", w)
	w.balance += amount
}

func (w *Wallet) Balance() int {
	return w.balance
}
