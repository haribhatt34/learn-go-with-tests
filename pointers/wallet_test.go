package pointers

import (
	"testing"
)

func TestWallet(t *testing.T) {
	wallet := Wallet{}
	wallet.Deposit(Bitcoin(10))

	balance := wallet.Balance()
	want := Bitcoin(10)

	if balance != want {
		t.Errorf("got %d want %d", balance, want)
	}
}
