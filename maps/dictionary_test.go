package maps_test

import (
	"errors"
	"testing"

	"learn-go-with-tests/maps"
)

func TestSearch(t *testing.T) {
	t.Parallel()

	dictionary := maps.Dictionary{"test": "this is just a test"}

	t.Run("Search finds known word", func(t *testing.T) {
		t.Parallel()

		got, _ := dictionary.Search("test")
		want := "this is just a test"

		assertStrings(t, got, want)
	})

	t.Run("Search throws error for unknown word", func(t *testing.T) {
		t.Parallel()

		_, err := dictionary.Search("unknown")

		assertError(t, err, maps.ErrWordNotFound)
	})
}

func TestAdd(t *testing.T) {
	t.Parallel()

	t.Run("Add adds a new word", func(t *testing.T) {
		t.Parallel()
		dictionary := maps.Dictionary{}
		word := "TDD"
		definition := "Test-driven development"

		err := dictionary.Add(word, definition)

		assertError(t, err, nil)
		assertDefinition(t, dictionary, word, definition)
	})

	t.Run("Add throws an error for existing word", func(t *testing.T) {
		t.Parallel()
		word := "test"
		definition := "this is just a test"
		dictionary := maps.Dictionary{word: definition}

		err := dictionary.Add(word, "new test")

		assertError(t, err, maps.ErrWordExists)
		assertDefinition(t, dictionary, word, definition)
	})
}

func TestUpdate(t *testing.T) {
	t.Parallel()

	t.Run("Update works with existing word", func(t *testing.T) {
		t.Parallel()
		word := "test"
		definition := "this is an update test"
		dictionary := maps.Dictionary{word: definition}
		newDefinition := "updated"

		err := dictionary.Update(word, newDefinition)

		assertError(t, err, nil)
		assertDefinition(t, dictionary, word, newDefinition)
	})

	t.Run("Update fails with new word", func(t *testing.T) {
		t.Parallel()
		word := "newWord"
		dictionary := maps.Dictionary{}
		newDefinition := "updated"

		err := dictionary.Update(word, newDefinition)

		assertError(t, err, maps.ErrWordDoesNotExist)
	})
}

func TestDelete(t *testing.T) {
	t.Parallel()

	word := "toBeDeleted"
	dictionary := maps.Dictionary{word: "test definition"}

	dictionary.Delete(word)

	_, err := dictionary.Search(word)
	if !errors.Is(err, maps.ErrWordNotFound) {
		t.Errorf("Expected %q to be deleted", word)
	}
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

func assertDefinition(
	tb testing.TB,
	dictionary maps.Dictionary,
	word,
	definition string,
) {
	tb.Helper()

	got, err := dictionary.Search(word)
	if err != nil {
		tb.Fatal("should find added word:", err)
	}

	assertStrings(tb, got, definition)
}
