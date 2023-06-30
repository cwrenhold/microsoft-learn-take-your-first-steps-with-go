package main

import (
	"fmt"
	"log"
)

func main() {
	log.Panic("Hey, I'm an error log!")

	// Note: This line will not be executed as the log.Panic() call
	//       above will terminate the program. In addition to this,
	//       the log.Panic() call will also print the error message
	//       and stack trace to the console.
	fmt.Print("Can you see me?")
}
