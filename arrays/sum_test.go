package array_test

import (
	"testing"

	array "learn-go-with-tests/arrays"
)

func TestSum(t *testing.T) {
	t.Parallel()

	numbers := [5]int{1, 2, 3, 4, 5}

	got := array.Sum(numbers)
	want := 15

	if got != want {
		t.Errorf("got %d want %d given %v", got, want, numbers)
	}
}
