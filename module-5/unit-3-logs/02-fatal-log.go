package main

import (
	"fmt"
	"log"
)

func main() {
	log.Fatal("Hey, I'm an error log!")

	// Note: This line will not be executed as the log.Fatal() call
	//       above will terminate the program.
	fmt.Print("Can you see me?")
}
