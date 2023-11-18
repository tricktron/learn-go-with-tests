package pointers_test

import (
	"errors"
	"testing"

	"learn-go-with-tests/pointers"
)

func TestWallet(t *testing.T) {
	t.Parallel()

	t.Run("Wallet deposits a positive amount correctly", func(t *testing.T) {
		t.Parallel()

		wallet := pointers.Wallet{}

		wallet.Deposit(pointers.Bitcoin(10))

		assertBalance(t, wallet, pointers.Bitcoin(10))
	})

	t.Run("Wallet withdraws a positive amount correctly", func(t *testing.T) {
		t.Parallel()

		wallet := pointers.NewWallet(pointers.Bitcoin(20))

		wallet.Withdraw(pointers.Bitcoin(10)) //nolint: errcheck

		assertBalance(t, *wallet, pointers.Bitcoin(10))
	})

	t.Run("Wallet cannot withdraw a bigger amount than its balance",
		func(t *testing.T) {
			t.Parallel()

			startingBalance := pointers.Bitcoin(20)
			wallet := pointers.NewWallet(startingBalance)

			err := wallet.Withdraw(pointers.Bitcoin(100))

			assertBalance(t, *wallet, startingBalance)
			assertError(t, err, pointers.ErrInsufficientFunds)
		})
}

func assertBalance(
	tb testing.TB,
	wallet pointers.Wallet,
	want pointers.Bitcoin,
) {
	tb.Helper()

	got := wallet.Balance()

	if got != want {
		tb.Errorf("got %s want %s", got, want)
	}
}

func assertError(tb testing.TB, got, want error) {
	tb.Helper()

	if got == nil {
		tb.Fatal("wanted an error but did not get one")
	}

	if !errors.Is(got, want) {
		tb.Errorf("got %q want %q", got, want)
	}
}
