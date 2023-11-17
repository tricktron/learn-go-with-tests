package iteration

// Repeats the character iterations times.
func Repeat(character string, iterations int) string {
	var repeated string
	for i := 0; i < iterations; i++ {
		repeated += character
	}

	return repeated
}
