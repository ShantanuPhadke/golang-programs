package shapes

import "math"

// CHANGE LOG:
// (1) Added a function to calculate Perimeter of a rectangle.
// (2) Added a function to calculate the Area of a rectangle.
// (3) Define a type of Rectangle with the attributes width and height
// (4) Make a type of Circle with attribute radius
// (5) Define area methods in our two types, Circle and Rectangle
// (6) Defined a general Shape interface to avoid code duplication

/*Shape ...
Methods:
	Area = Calculates area of given shape
*/
type Shape interface {
	Area() float64
}

/*Rectangle ...
ATTRIBUTES:
	width = width of the rectangle represented
	height = height of the rectangle represented
*/
type Rectangle struct {
	width  float64
	height float64
}

/*Area ... (For Rectangles) */
func (r Rectangle) Area() float64 {
	return r.width * r.height
}

/*Circle ...
radius = radius of the circle
*/
type Circle struct {
	radius float64
}

/*Area ... (For Circles) */
func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

/*Triangle ...
base = base of the triangle
height = height of the triangle
*/
type Triangle struct {
	base   float64
	height float64
}

/*Area ... (For Triangle) */
func (t Triangle) Area() float64 {
	return 0.4 * t.base * t.height
}

/*Perimeter ...
INPUTS:
	width = width of the rectangle whose perimeter we want to calculate
	height = height of the rectangle whose perimeter we want to calculate
OUTPUTS:
	Returns the perimeter of the rectangle described by the input dimensions.
*/
func Perimeter(rectangle Rectangle) float64 {
	return 2*rectangle.width + 2*rectangle.height
}

/*Area ...
INPUTS:
	width = width of the rectangle whose area we want to calculate
	height = height of the rectangle whose area we want to calculate
OUTPUTS:
	Returns the area of the rectangle described by the input dimensions
*/
func Area(rectangle Rectangle) float64 {
	return rectangle.width * rectangle.height
}
