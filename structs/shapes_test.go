package structs

import (
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

	/* By using interface, we have decoupled the helper method from
	 * the concrete types (rectange, circle), now we can pass any
	 * type which is a shape and get the job done.
	 * In an essence code is also refactored, we don't have to write
	 * assert case for every comparison.
	 * If an object is passed which is not shaped, it will safely fail.
	 */

	/* Change to table based test below
	checkArea := func(t *testing.T, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()
		if got != want {
			t.Errorf("got %g want %g", got, want)
		}
	}

	t.Run("rectangles", func(t *testing.T) {
		rectangle := Rectangle{2.0, 3.0}
		want := 6.0
		checkArea(t, rectangle, want)

	})

	t.Run("circles", func(t *testing.T) {
		circle := Circle{10.0}
		want := 314.15926
		checkArea(t, circle, want)
	})
	*/

	areaTests := []struct {
		name    string
		shape   Shape
		hasArea float64
	}{
		{name: "Rectangle", shape: Rectangle{2.0, 3.0}, hasArea: 6.0},
		{name: "Circle", shape: Circle{10.0}, hasArea: 314.1592653589793},
		{name: "Triangle", shape: Triangle{12, 6}, hasArea: 36.0},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.hasArea {
				t.Errorf("%#v got %g want %g", tt.shape, got, tt.hasArea)
			}
		})
	}
}

/* Quote:
 * In Test-Driven Deveopment by Example Kent Beck refactors some tests to
 * a point and asserts:
 * The test speaks to us more clearly, as if it were an assertion of truth,
 * not a sequenc of operations.
 */

/* Notes:
 * Declaring interface helps in defining functions that can be used by
 * different types (parameteric polymorphism)
 * Adding methods so we can add functionality to our data types and also
 * we can implement interfaces.
 * Table based test like above, make assertions clearer and easier to
 * extend & maintain.
 */

/* Important:
 * Interfaces are a great way to hide complexity.
 * In our case test helper didn't need to know the exact shape it was
 * asserting on, it only need to know the it area.
 */
