package main

import "fmt"

func main() {
	studentsAge := make(map[string]int)
	studentsAge["john"] = 32
	studentsAge["bob"] = 31
	fmt.Println(studentsAge)

	// Note: We must us the make function to initialise a map before we can
	//       add items to it. If we defined the map using the following
	//       statement:
	//       var studentsAge map[string]int
	//       we would get a runtime error when we try to add items to it.
}
