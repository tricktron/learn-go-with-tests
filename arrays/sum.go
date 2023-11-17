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
	sumSlice := make([]int, 0, len(slices))

	for _, slice := range slices {
		sumSlice = append(sumSlice, Sum(slice))
	}

	return sumSlice
}

func SumAllTails(slices ...[]int) []int {
	sumSlice := make([]int, 0, len(slices))

	for _, slice := range slices {
		sumSlice = append(sumSlice, Sum(slice[1:]))
	}

	return sumSlice
}
