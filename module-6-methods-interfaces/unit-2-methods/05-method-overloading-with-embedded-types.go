package main

import "fmt"

type triangle struct {
	size int
}

func (t triangle) perimeter() int {
	return t.size * 3
}

type coloredTriangle struct {
	triangle
	color string
}

// Note: Even though the coloredTriangle class has an embedded
// type of triangle, it can still have its own perimeter method.
// This will then mean that the original perimeter method will
// be shadowed and inaccessible.
func (t coloredTriangle) perimeter() int {
	return t.size * 3 * 2
}

func main() {
	t := coloredTriangle{triangle{3}, "blue"}
	fmt.Println("Size:", t.size)
	fmt.Println("Perimeter (colored)", t.perimeter())
	fmt.Println("Perimeter (normal)", t.triangle.perimeter())
}
