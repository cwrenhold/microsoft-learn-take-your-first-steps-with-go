package main

import "fmt"

func main() {
	var numbers []int
	// Note: The capacity of the slice is doubled each time it is exceeded.
	for i := 0; i < 10; i++ {
		numbers = append(numbers, i)
		fmt.Printf("%d\tcap=%d\t%v\n", i, cap(numbers), numbers)
	}
}
