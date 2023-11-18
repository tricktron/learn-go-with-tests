package pointers_test

import (
	"testing"

	"learn-go-with-tests/pointers"
)

func TestWallet(t *testing.T) {
	t.Parallel()

	wallet := pointers.Wallet{}

	wallet.Deposit(pointers.Bitcoin(10))
	got := wallet.Balance()
	want := pointers.Bitcoin(99)

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}
