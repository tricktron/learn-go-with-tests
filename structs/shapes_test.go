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

	areaTests := []struct {
		name  string
		shape structs.Shape
		area  float64
	}{
		{
			name:  "Rectangle",
			shape: structs.Rectangle{Width: 5, Height: 7},
			area:  35.0,
		},
		{
			name:  "Circle",
			shape: structs.Circle{Radius: 10},
			area:  314.1592653589793,
		},
		{
			name:  "Triangle",
			shape: structs.Triangle{Base: 12, Height: 6},
			area:  36.0,
		},
	}

	for _, tc := range areaTests {
		tc := tc //nolint: varnamelen
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			got := tc.shape.Area()
			if got != tc.area {
				t.Errorf("%#v got %g want %g", tc.shape, got, tc.area)
			}
		})
	}
}
