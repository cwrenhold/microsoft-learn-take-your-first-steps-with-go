package main

import (
	"fmt"

	"github.com/cwrenhold/calculator"
)

func main() {
	total := calculator.Sum(3, 5)
	fmt.Println(total)
	fmt.Println("Version: ", calculator.Version)
}
