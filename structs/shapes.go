package structs

import "math"

// Shape is an interface for rectangel and circle
type Shape interface {
	Area() float64
}

// Rectangle defines the parameter of rectangle
type Rectangle struct {
	length  float64
	breadth float64
}

// Circle defines the parameter of Circle
type Circle struct {
	radius float64
}

// Triangle defines the parameter of Triangle
type Triangle struct {
	base   float64
	height float64
}

// Perimeter function returns perimeter of rectangle
func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.length + rectangle.breadth)
}

// Area function returns area of rectangle
func (r Rectangle) Area() float64 {
	return r.length * r.breadth
}

// Area function returns area of circle
func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

// Area function returns area of triangle
func (t Triangle) Area() float64 {
	return 0.5 * t.base * t.height
}
