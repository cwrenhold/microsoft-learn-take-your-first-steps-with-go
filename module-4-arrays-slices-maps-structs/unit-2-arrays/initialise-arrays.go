package main

import "fmt"

func main() {
	// Note: The array is of length 5, but we only assign 4 values, so
	//       the last one is the default value, i.e. an empty string
	cities := [5]string{"New York", "Paris", "Berlin", "Madrid"}
	fmt.Println("Cities:", cities)
}
