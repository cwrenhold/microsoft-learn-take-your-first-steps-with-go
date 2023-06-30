package main

import "fmt"

func main() {
	studentsAge := make(map[string]int)
	studentsAge["john"] = 32
	studentsAge["bob"] = 31
	delete(studentsAge, "john")
	fmt.Println(studentsAge)

	// Note: This doesn't cause an error if the key doesn't exist.
	delete(studentsAge, "christy")
}
