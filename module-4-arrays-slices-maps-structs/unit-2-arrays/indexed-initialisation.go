package main

import "fmt"

func main() {
	// Note: The array is defined with an ellipsis, but we are only
	//       assigning a value to position 99, so the first 98 values
	//       are the default value, i.e. 0
	numbers := [...]int{99: -1}
	fmt.Println("First Position:", numbers[0])
	fmt.Println("Last Position:", numbers[99])
	fmt.Println("Length:", len(numbers))
}
