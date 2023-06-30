package main

import "fmt"

type triangle struct {
	size int
}

func (t triangle) perimeter() int {
	return t.size * 3
}

// Note: This class gets the perimeter method from the triangle
// class because it has an embedded type of triangle. This means
// that the coloredTriangle class has access to all of the methods
// of the triangle class.
// This isn't inheritance, it's composition. The coloredTriangle
// class has a triangle, not is a triangle. So the permimeter()
// method is called on the triangle class, not the coloredTriangle,
// and effectively compiles to t.triangle.perimeter() where called.
type coloredTriangle struct {
	triangle
	color string
}

func main() {
	t := coloredTriangle{triangle{3}, "blue"}
	fmt.Println("Size:", t.size)
	fmt.Println("Perimeter", t.perimeter())
}
