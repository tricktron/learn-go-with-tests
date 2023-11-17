package structs

func Perimeter(height, width float64) float64 {
	return 2 * (height + width) //nolint: gomnd
}

func Area(height, width float64) float64 {
	return height * width
}
