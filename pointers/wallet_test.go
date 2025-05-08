package wallet

import (
	"fmt"
	"testing"
)

func TestWallet(t *testing.T) {

	assertBalance := func(t testing.TB, wallet Wallet, want Bitcoin) {
		t.Helper()
		got := wallet.Balance()

		if got != want {
			t.Errorf("wanted %s but got %s", want, got)
		}
	}

	assertError := func(t testing.TB, err error) {
		t.Helper()

		if err == nil {
			t.Fatal("wanted error but didnt get one")
		}

		if err.Error() != ErrInsufficientFunds {
			t.Errorf("want %s but got %s", ErrInsufficientFunds, err.Error())
		}
	}

	assertNoError := func(t testing.TB, err error) {
		t.Helper()

		if err != nil {
			t.Errorf("wanted no error but got one")
		}
	}

	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))

		fmt.Printf("address of balance in test is %p\n", &wallet)

		want := Bitcoin(10)

		assertBalance(t, wallet, want)
	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		err := wallet.Withdraw(Bitcoin(10))
		want := Bitcoin(10)

		assertBalance(t, wallet, want)
		assertNoError(t, err)
	})

	t.Run("Insufficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(10)
		wallet := Wallet{startingBalance}
		err := wallet.Withdraw(Bitcoin(100))

		assertBalance(t, wallet, startingBalance)
		assertError(t, err)
	})
}
