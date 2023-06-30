package main

import "fmt"

func main() {
	value := 0
	for value <= 0 {
		fmt.Print("Number of elements to calculate: ")
		fmt.Scanf("%d", &value)

		if value <= 0 {
			fmt.Println("Please enter a positive number")
		}
	}

	fib := fibonacci(value)

	fmt.Println(fib)
}

func fibonacci(max int) []int {
	// Add a quick exit if the max is less than 2.
	if max < 2 {
		return make([]int, 0)
	}

	// Declare the slice to store the Fibonacci numbers.
	fib := make([]int, max)
	fib[0] = 1
	fib[1] = 1

	// Iterate over the remaining numbers and append the Fibonacci numbers to the slice.
	for i := 2; i < max; i++ {
		fib[i] = fib[i-1] + fib[i-2]
	}

	return fib
}
