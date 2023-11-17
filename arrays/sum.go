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
		if len(slice) == 0 {
			sumSlice = append(sumSlice, 0)
		} else {
			sumSlice = append(sumSlice, Sum(slice[1:]))
		}
	}

	return sumSlice
}
