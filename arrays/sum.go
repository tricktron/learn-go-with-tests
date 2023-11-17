package array

func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}

	return sum
}

// Creates a new slice with the sum of slices.
func SumAll(slices ...[]int) []int {
	slicesCount := len(slices)
	sumArray := make([]int, slicesCount)

	for i, slice := range slices {
		sumArray[i] = Sum(slice)
	}

	return sumArray
}
