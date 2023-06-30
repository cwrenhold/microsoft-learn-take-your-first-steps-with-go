package main

import "fmt"

// Note: The channel here is a send-only channel.
//
//	This is indicated by the chan<- syntax.
func send(ch chan<- string, message string) {
	fmt.Printf("Sending: %#v\n", message)
	ch <- message

	// Note: This would be a compile-time error as we are trying to receive
	//       from a send-only channel.
	//<-ch
}

// Note: The channel here is a receive-only channel.
//
//	This is indicated by the <-chan syntax.
func read(ch <-chan string) {
	fmt.Printf("Receiving: %#v\n", <-ch)
}

func main() {
	ch := make(chan string, 1)
	send(ch, "Hello World!")
	read(ch)
}
