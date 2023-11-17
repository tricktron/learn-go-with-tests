package structs_test

import (
	"testing"

	"learn-go-with-tests/structs"
)

func TestPerimeter(t *testing.T) {
	t.Parallel()

	got := structs.Perimeter(10.0, 10.0)
	want := 40.0

	if got != want {
		t.Errorf("got %.2f want %2.f", got, want)
	}
}
