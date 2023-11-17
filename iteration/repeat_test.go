package iteration_test

import (
	"testing"

	"learn-go-with-tests/iteration"
)

func TestRepeat(t *testing.T) {
	t.Parallel()

	got := iteration.Repeat("a")
	want := "aaaaa"

	if got != want {
		t.Errorf("want %q but got %q", want, got)
	}
}
