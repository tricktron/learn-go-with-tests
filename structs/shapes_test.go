package structs_test

import (
	"testing"

	"learn-go-with-tests/structs"
)

func TestPerimeter(t *testing.T) {
	t.Parallel()

	rectangle := structs.Rectangle{10.0, 10.0}
	got := rectangle.Perimeter()
	want := 40.0

	if got != want {
		t.Errorf("got %.2f want %2.f", got, want)
	}
}

func TestArea(t *testing.T) {
	t.Parallel()

	t.Run("Should calculate area of rectangles", func(t *testing.T) {
		t.Parallel()

		rectangle := structs.Rectangle{5.0, 7.0}
		got := rectangle.Area()
		want := 35.0

		if got != want {
			t.Errorf("got %.2f want %2.f", got, want)
		}
	})

	t.Run("Should calculate area of circles", func(t *testing.T) {
		t.Parallel()

		circle := structs.Circle{10}
		got := circle.Area()
		want := 314.1592653589793

		if got != want {
			t.Errorf("got %g want %g", got, want)
		}
	})
}
