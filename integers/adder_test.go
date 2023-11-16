package integers_test

import (
	"fmt"
	"testing"

	"learn-go-with-tests/integers"
)

func TestAdder(t *testing.T) {
	t.Parallel()

	got := integers.Add(2, 2)
	want := 4

	if got != want {
		t.Errorf("expected '%d' but got '%d'", want, got)
	}
}

func ExampleAdd() {
	sum := integers.Add(38, 4)
	fmt.Println(sum)
	// Output: 42
}
