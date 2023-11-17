package iteration_test

import (
	"fmt"
	"testing"

	"learn-go-with-tests/iteration"
)

func TestRepeat(t *testing.T) {
	t.Parallel()

	got := iteration.Repeat("a", 5)
	want := "aaaaa"

	if got != want {
		t.Errorf("want %q but got %q", want, got)
	}
}

var result string //nolint: gochecknoglobals

func BenchmarkRepeat(b *testing.B) {
	var repeated string
	for i := 0; i < b.N; i++ {
		// always record the result to prevent the compiler eliminating the
		// function call
		repeated = iteration.Repeat("a", 5)
	}

	// always store the result to a package level variable so that the compiler
	// cannot eliminate the benchmark itself
	result = repeated
}

func ExampleRepeat() {
	repeated := iteration.Repeat("x", 3)
	fmt.Println(repeated)
	// Output: xxx
}
