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

	assertError := func(tb testing.TB, got error, want string) {
		tb.Helper()

		if got == nil {
			t.Fatal("wanted an error but did not get one")
		}

		if got.Error() != want {
			t.Errorf("got %q want %q", got, want)
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
			assertError(t, err, "cannot withdraw, insufficient funds")
		})
}
