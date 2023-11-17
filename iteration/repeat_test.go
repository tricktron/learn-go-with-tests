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

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		iteration.Repeat("a", 5)
	}
}

func ExampleRepeat() {
	repeated := iteration.Repeat("x", 3)
	fmt.Println(repeated)
	// Output: xxx
}
