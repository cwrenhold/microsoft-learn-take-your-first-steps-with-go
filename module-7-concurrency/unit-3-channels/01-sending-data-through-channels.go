package main

import (
	"fmt"
	"net/http"
	"time"
)

func checkAPI(api string, ch chan string) {
	_, err := http.Get(api)
	if err != nil {
		ch <- fmt.Sprintf("ERROR: %s is down!\n", api)
		return
	}

	ch <- fmt.Sprintf("SUCCESS: %s is up and running!\n", api)
}

func main() {
	start := time.Now()

	apis := []string{
		"https://management.azure.com",
		"https://dev.azure.com",
		"https://api.github.com",
		"https://outlook.office.com/",
		"https://api.somewhereintheinternet.com/",
		"https://graph.microsoft.com",
	}

	// Create a channel which allows strings
	ch := make(chan string)

	for _, api := range apis {
		go checkAPI(api, ch)
	}

	// Print the data received from the channel, whatever it is
	// Note: This is a blocking call, but only waits for the first data
	//       to be received, not all of them. This means that the first
	//       data to be received will be printed, and the rest will be
	//       ignored, and blocked until the first data is received.
	fmt.Print(<-ch)

	elapsed := time.Since(start)
	fmt.Printf("Done! It took %v seconds!\n", elapsed.Seconds())
}
