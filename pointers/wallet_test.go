package pointers_test

import (
	"testing"

	"learn-go-with-tests/pointers"
)

func TestWallet(t *testing.T) {
	t.Parallel()

	assertBalance := func(
		tb testing.TB,
		wallet pointers.Wallet,
		want pointers.Bitcoin,
	) {
		tb.Helper()

		got := wallet.Balance()

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	}

	t.Run("Wallet deposits a positive amount correctly", func(t *testing.T) {
		t.Parallel()

		wallet := pointers.Wallet{}

		wallet.Deposit(pointers.Bitcoin(10))

		assertBalance(t, wallet, pointers.Bitcoin(10))
	})

	t.Run("Wallet withdraws a positive amount correctly", func(t *testing.T) {
		t.Parallel()

		wallet := pointers.NewWallet(pointers.Bitcoin(20))

		wallet.Withdraw(pointers.Bitcoin(10))

		assertBalance(t, *wallet, pointers.Bitcoin(10))
	})
}
