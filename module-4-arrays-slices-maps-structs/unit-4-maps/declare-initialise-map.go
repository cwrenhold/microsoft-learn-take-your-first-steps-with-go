package main

import "fmt"

func main() {
	// Note: The key is the type inside the square brackets and the type of
	//       the value is after the closing square bracket.
	//       The items in the map are then initialised using curly brackets.
	studentsAge := map[string]int{
		"john": 32,
		"bob":  31,
	}
	fmt.Println(studentsAge)

	// You can also declare a map without initialising it by using the make
	// function. The make function takes the type of the map as a parameter.
	// E.g. make(map[string]int) would create a map with string keys and int
	// values.
}
