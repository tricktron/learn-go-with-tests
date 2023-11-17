package array_test

import (
	"testing"

	array "learn-go-with-tests/arrays"
)

func TestSum(t *testing.T) {
	t.Parallel()

	t.Run("Should calculate sum of an array with size 5", func(t *testing.T) {
		t.Parallel()

		numbers := []int{1, 2, 3, 4, 5}

		got := array.Sum(numbers)
		want := 15

		if got != want {
			t.Errorf("got %d want %d given %v", got, want, numbers)
		}
	})

	t.Run("Should calculate sum of an array with size 4", func(t *testing.T) {
		t.Parallel()

		numbers := []int{1, 2, 3, 4}

		got := array.Sum(numbers)
		want := 10

		if got != want {
			t.Errorf("got %d want %d given %v", got, want, numbers)
		}
	})
}
