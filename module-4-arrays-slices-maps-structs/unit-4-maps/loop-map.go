package main

import (
	"fmt"
)

func main() {
	studentsAge := make(map[string]int)
	studentsAge["john"] = 32
	studentsAge["bob"] = 31

	// Note: We can declare a variable for both the name and the age.
	//       These will then be set to the key and value for each item
	//       in the map.
	for name, age := range studentsAge {
		fmt.Printf("%s\t%d\n", name, age)
	}

	// Note: We can also discard the key using the underscore character.
	for _, age := range studentsAge {
		fmt.Printf("Ages %d\n", age)
	}

	// Note: If we only want the keys, we only need to use one variable.
	for name := range studentsAge {
		fmt.Printf("Names %s\n", name)
	}
}
