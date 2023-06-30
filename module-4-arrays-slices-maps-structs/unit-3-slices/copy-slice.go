package main

import "fmt"

func main() {
	letters := []string{"A", "B", "C", "D", "E"}
	fmt.Println("Before", letters)

	slice1 := letters[0:2]

	// Two step process to copy the values from the array into a new slice.
	// First we create a new slice with make, then we copy the values from
	// the array into the new slice.
	slice2 := make([]string, 3)
	copy(slice2, letters[1:4])

	slice1[1] = "Z"

	fmt.Println("After", letters)

	// Note: slice2 is not affected by the change to slice1.
	fmt.Println("Slice2", slice2)
}
