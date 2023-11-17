package structs_test

import (
	"testing"

	"learn-go-with-tests/structs"
)

func TestPerimeter(t *testing.T) {
	t.Parallel()

	rectangle := structs.Rectangle{10.0, 10.0}
	got := structs.Perimeter(rectangle)
	want := 40.0

	if got != want {
		t.Errorf("got %.2f want %2.f", got, want)
	}
}

func TestArea(t *testing.T) {
	t.Parallel()

	rectangle := structs.Rectangle{5.0, 7.0}
	got := structs.Area(rectangle)
	want := 35.0

	if got != want {
		t.Errorf("got %.2f want %2.f", got, want)
	}
}
