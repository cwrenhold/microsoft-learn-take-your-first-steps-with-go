package main

import (
	"log"
)

func main() {
	// Note: This adds a prefix to all log messages going forward.
	log.SetPrefix("main(): ")
	log.Print("Hey, I'm a log!")
	log.Fatal("Hey, I'm an error log!")
}
