package maps_test

import (
	"errors"
	"testing"

	"learn-go-with-tests/maps"
)

func TestSearch(t *testing.T) {
	t.Parallel()

	dictionary := maps.Dictionary{"test": "this is just a test"}

	t.Run("Dictionary finds known word", func(t *testing.T) {
		t.Parallel()

		got, _ := dictionary.Search("test")
		want := "this is just a test"

		assertStrings(t, got, want)
	})

	t.Run("Dictionary throws error for unknown word", func(t *testing.T) {
		t.Parallel()

		_, err := dictionary.Search("unknown")

		assertError(t, err, maps.ErrWordNotFound)
	})
}

func assertStrings(tb testing.TB, got, want string) {
	tb.Helper()

	if got != want {
		tb.Errorf("got %q want %q", got, want)
	}
}

func assertError(tb testing.TB, got, want error) {
	tb.Helper()

	if !errors.Is(got, want) {
		tb.Errorf("got error %q want %q", got, want)
	}
}
