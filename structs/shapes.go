package structs

import "math"

//Rectangle defines the parameter of rectangle
type Rectangle struct {
	length  float64
	breadth float64
}

//Circle defines the parameter of rectangle
type Circle struct {
	radius float64
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
