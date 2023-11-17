package array_test

import (
	"reflect"
	"testing"

	array "learn-go-with-tests/arrays"
)

func TestSum(t *testing.T) {
	t.Parallel()

	t.Run("Should calculate sum of non-empty slice", func(t *testing.T) {
		t.Parallel()

		numbers := []int{1, 2, 3, 4, 5}

		got := array.Sum(numbers)
		want := 15

		if got != want {
			t.Errorf("got %d want %d given %v", got, want, numbers)
		}
	})

	t.Run("Should return 0 sum for empty slice", func(t *testing.T) {
		t.Parallel()

		numbers := []int{}

		got := array.Sum(numbers)
		want := 0

		if got != want {
			t.Errorf("got %d want %d given %v", got, want, numbers)
		}
	})
}

func TestSumAll(t *testing.T) {
	t.Parallel()

	t.Run("Should calculate and return slices sum of single slice",
		func(t *testing.T) {
			t.Parallel()

			slice1 := []int{1, 2, 3, 4, 5}

			got := array.SumAll(slice1)
			want := []int{15}

			if !reflect.DeepEqual(got, want) {
				t.Errorf("got %v want %v", got, want)
			}
		})

	t.Run("Should calculate and return slices sum of multiple slices",
		func(t *testing.T) {
			t.Parallel()

			slice1 := []int{1, 2}
			slice2 := []int{41, 1}

			got := array.SumAll(slice1, slice2)
			want := []int{3, 42}

			if !reflect.DeepEqual(got, want) {
				t.Errorf("got %v want %v", got, want)
			}
		})
}

func TestSumAllTails(t *testing.T) {
	t.Parallel()

	t.Run("Should calculate and return slices tail sum of single slice",
		func(t *testing.T) {
			t.Parallel()

			slice1 := []int{1, 2, 3, 4, 5}

			got := array.SumAllTails(slice1)
			want := []int{14}

			if !reflect.DeepEqual(got, want) {
				t.Errorf("got %v want %v", got, want)
			}
		})

	t.Run("Should calculate and return slices tail sum of multiple slices",
		func(t *testing.T) {
			t.Parallel()

			slice1 := []int{1, 2}
			slice2 := []int{41, 1}

			got := array.SumAllTails(slice1, slice2)
			want := []int{2, 1}

			if !reflect.DeepEqual(got, want) {
				t.Errorf("got %v want %v", got, want)
			}
		})
}
