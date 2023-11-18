package pointers_test

import (
	"testing"

	"learn-go-with-tests/pointers"
)

func TestWallet(t *testing.T) {
	t.Parallel()

	wallet := pointers.Wallet{}

	wallet.Deposit(10)
	got := wallet.Balance()
	want := 10

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
