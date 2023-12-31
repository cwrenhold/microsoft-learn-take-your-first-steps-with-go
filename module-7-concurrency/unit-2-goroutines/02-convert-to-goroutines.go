package main

import (
	"fmt"
	"net/http"
	"time"
)

func checkAPI(api string) {
	_, err := http.Get(api)
	if err != nil {
		fmt.Printf("ERROR: %s is down!\n", api)
		return
	}

	fmt.Printf("SUCCESS: %s is up and running!\n", api)
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

	for _, api := range apis {
		// The go keyword is used to create a goroutine
		go checkAPI(api)
	}

	// Sleep so that the goroutines can finish
	time.Sleep(time.Second * 3)

	// Note: This isn't realiable now because of the sleep
	elapsed := time.Since(start)
	fmt.Printf("Done! It took %v seconds!\n", elapsed.Seconds())
}
