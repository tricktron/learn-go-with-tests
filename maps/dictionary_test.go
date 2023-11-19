package maps_test

import (
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
		want := "could not find the word you were looking for"

		if err == nil {
			t.Fatal("expected to get an error")
		}

		assertStrings(t, err.Error(), want)
	})
}

func assertStrings(tb testing.TB, got, want string) {
	tb.Helper()

	if got != want {
		tb.Errorf("got %q want %q", got, want)
	}
}
