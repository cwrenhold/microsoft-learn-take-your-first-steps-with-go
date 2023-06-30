package main

import "fmt"

func main() {
	studentsAge := make(map[string]int)
	studentsAge["john"] = 32
	studentsAge["bob"] = 31
	fmt.Println("Bob's age is", studentsAge["bob"])

	// Note: If we use a key that doesn't exist in the map, we get the default value.
	fmt.Println("Christy's age is", studentsAge["christy"])
}
