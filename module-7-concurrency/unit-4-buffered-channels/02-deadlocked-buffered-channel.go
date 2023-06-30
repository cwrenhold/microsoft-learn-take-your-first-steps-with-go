package main

import (
	"fmt"
)

func send(ch chan string, message string) {
	ch <- message
}

func main() {
	// With this size, the channel isn't big enough to hold all the data
	// we are sending to it. This will cause a deadlock as we don't take
	// any data out of the channel until all items have been sent to it,
	// and we can't send all items to it because it's full.
	// This could be solved by making the calls to send() concurrent.
	size := 2
	ch := make(chan string, size)
	send(ch, "one")
	send(ch, "two")
	send(ch, "three")
	send(ch, "four")
	fmt.Println("All data sent to the channel ...")

	for i := 0; i < size; i++ {
		fmt.Println(<-ch)
	}

	fmt.Println("Done!")
}
