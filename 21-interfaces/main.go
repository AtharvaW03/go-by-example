package main

import (
	"fmt"
	"math"
)

// Interface (contract).
// Any type that provides area() and perimeter()
// automatically satisfies geometry.
type geometry interface {
	area() float64
	perimeter() float64
}

// Rectangle type.
type rect struct {
	width, height float64
}

// Circle type.
type circle struct {
	radius float64
}

// rect satisfies geometry by implementing area().
func (r rect) area() float64 {
	return r.width * r.height
}

// rect satisfies geometry by implementing perimeter().
func (r rect) perimeter() float64 {
	return 2*r.width + 2*r.height
}

// circle satisfies geometry by implementing area().
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

// circle satisfies geometry by implementing perimeter().
func (c circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

// Accepts anything that satisfies the geometry interface.
// Doesn't care whether it's a rect, circle, etc.
func measure(g geometry) {
	fmt.Println(g)

	// Calls the appropriate implementation
	// based on the actual type stored in g.
	fmt.Println(g.area())
	fmt.Println(g.perimeter())
}

// Demonstrates a type assertion.
func detectCircle(g geometry) {

	// Ask:
	// "Is the concrete type inside g actually a circle?"
	//
	// If yes:
	//   c  = the circle value
	//   ok = true
	//
	// If no:
	//   ok = false
	if c, ok := g.(circle); ok {
		fmt.Println("circle with radius", c.radius)
	}
}

func main() {

	// Create a rectangle.
	r := rect{width: 3, height: 4}
	// Create a circle.
	c := circle{radius: 5}

	// Both rect and circle satisfy geometry,
	// so both can be passed to measure().
	measure(r)
	measure(c)

	// r is not a circle, so nothing is printed.
	detectCircle(r)
	// c is a circle, so the type assertion succeeds.
	detectCircle(c)
}
