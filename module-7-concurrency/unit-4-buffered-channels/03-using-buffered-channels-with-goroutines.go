package main

import (
	"fmt"
)

func send(ch chan string, message string) {
	ch <- message
}

func main() {
	size := 2
	ch := make(chan string, size)
	send(ch, "one")
	send(ch, "two")
	// With these two calls sent as goroutines, the channel is big enough
	// to hold all the data we are sending to it. This will not cause a
	// deadlock as these requests will be locked until the first items are
	// taken out of the channel.
	go send(ch, "three")
	go send(ch, "four")
	fmt.Println("All data sent to the channel ...")

	// Note: We've changed this to a loop of 4 iterations, as we aren't
	//       tied to the size of the channel anymore.
	for i := 0; i < 4; i++ {
		fmt.Println(<-ch)
	}

	fmt.Println("Done!")
}
