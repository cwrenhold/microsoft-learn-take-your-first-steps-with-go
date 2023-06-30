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

func (s square) perimeter() int {
	return s.size * 4
}

func main() {
	t := triangle{3}
	s := square{4}
	fmt.Println("Perimeter (triangle):", t.perimeter())
	fmt.Println("Perimeter (square):", s.perimeter())
}
