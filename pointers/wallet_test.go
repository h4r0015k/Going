package wallet

import (
	"fmt"
	"testing"
)

func TestWallet(t *testing.T) {
	wallet := Wallet{}

	wallet.Deposit(10)

	got := wallet.Balance()
	fmt.Printf("address of balance in test is %p\n", &wallet)

	want := 10

	if got != want {
		t.Errorf("wanted %d but got %d", want, got)
	}
}
