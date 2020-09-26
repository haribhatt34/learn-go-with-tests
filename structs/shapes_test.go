package structs

import (
	"math"
	"testing"
)

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{
		length:  2.0,
		breadth: 3.0,
	}
	got := Perimeter(rectangle)
	want := 10.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {

	t.Run("rectangles", func(t *testing.T) {
		rectangle := Rectangle{2.0, 3.0}
		got := rectangle.Area()
		want := 6.0

		if got != want {
			t.Errorf("got %g want %.2f", got, want)
		}
	})

	t.Run("circles", func(t *testing.T) {
		circle := Circle{2.0}
		got := circle.Area()
		want := 4 * math.Pi

		if got != want {
			t.Errorf("got %g want %.2f", got, want)
		}
	})
}
