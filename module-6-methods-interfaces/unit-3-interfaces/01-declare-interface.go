package main

import "fmt"

type Shape interface {
	Perimeter() float64
	Area() float64
}

// Note: Square implicitly implements the Shape interface.
type Square struct {
	size float64
}

func (s Square) Area() float64 {
	return s.size * s.size
}

func (s Square) Perimeter() float64 {
	return s.size * 4
}

func main() {
	// Note: We're using the interface type Shape to hold a Square value.
	// We're explicitly declaring the type of the variable s to be Shape,
	// and then assigning a Square value to it.
	var s Shape = Square{3}
	fmt.Printf("%T\n", s)
	fmt.Println("Area: ", s.Area())
	fmt.Println("Perimeter:", s.Perimeter())
}
