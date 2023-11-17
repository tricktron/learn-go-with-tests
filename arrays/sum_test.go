package array_test

import (
	"testing"

	array "learn-go-with-tests/arrays"
)

func TestSum(t *testing.T) {
	t.Parallel()

	t.Run("Should calculate sum of non-empty array", func(t *testing.T) {
		t.Parallel()

		numbers := []int{1, 2, 3, 4, 5}

		got := array.Sum(numbers)
		want := 15

		if got != want {
			t.Errorf("got %d want %d given %v", got, want, numbers)
		}
	})

	t.Run("Should return 0 sum for empty array", func(t *testing.T) {
		t.Parallel()

		numbers := []int{}

		got := array.Sum(numbers)
		want := 0

		if got != want {
			t.Errorf("got %d want %d given %v", got, want, numbers)
		}
	})
}
