package main

import "fmt"

func main() {
	// Note: The array is of length 4 as we are using the ellipsis and
	//       have supplied 4 values.
	cities := [...]string{"New York", "Paris", "Berlin", "Madrid"}
	fmt.Println("Cities:", cities)
}
