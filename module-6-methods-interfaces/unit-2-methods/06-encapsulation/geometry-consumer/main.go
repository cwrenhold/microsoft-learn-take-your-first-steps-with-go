package main

import (
	"fmt"

	"github.com/cwrenhold/geometry"
)

func main() {
	t := geometry.Triangle{}
	t.SetSize(3)
	fmt.Println("Perimeter", t.Perimeter())

	// Note: Even though the type is exported, the size field is not.
	// fmt.Println("Size", t.size)
}
