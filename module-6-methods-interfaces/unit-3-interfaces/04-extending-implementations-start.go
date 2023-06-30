package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	// Make an HTTP GET request to the GitHub API for the Microsoft repos
	resp, err := http.Get("https://api.github.com/users/microsoft/repos?page=15&per_page=5")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	// Copy the response body to stdout
	io.Copy(os.Stdout, resp.Body)
}
