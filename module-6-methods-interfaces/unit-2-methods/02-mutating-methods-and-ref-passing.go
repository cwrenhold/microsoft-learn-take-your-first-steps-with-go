package main

import "fmt"

type triangle struct {
	size int
}

type square struct {
	size int
}

// Note: The receiver is in brackets and is the first parameter.
//
//	The receiver is a *copy* of the struct, so if you want to
//	modify the struct, you need to use a pointer.
func (t triangle) perimeter() int {
	return t.size * 3
}

// Note: The type is a pointer to a triangle. This is because we
// want to modify the struct, not just a copy of it.
// Note: It is convention to make all methods for a type either
// value or pointer receivers. This is because it can be confusing
// to have some methods that modify the struct and some that don't.
// This means that perimeter() should be a pointer receiver too.
func (t *triangle) doubleSize() {
	t.size *= 2
}

func (s square) perimeter() int {
	return s.size * 4
}

func main() {
	t := triangle{3}
	s := square{4}
	fmt.Println("Perimeter (triangle):", t.perimeter())
	fmt.Println("Perimeter (square):", s.perimeter())

	t.doubleSize()
	fmt.Println("Size (triangle):", t.size)
	fmt.Println("Perimeter (triangle):", t.perimeter())
}
