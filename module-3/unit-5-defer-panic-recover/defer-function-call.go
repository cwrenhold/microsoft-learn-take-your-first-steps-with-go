package main

import "fmt"

func doTheThing(i int) {
	defer fmt.Println("deferred", -i)
	fmt.Println("regular", i)
}

func main() {
	for i := 1; i <= 4; i++ {
		doTheThing(i)
	}
}
