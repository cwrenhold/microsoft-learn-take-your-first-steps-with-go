package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type GitHubResponse []struct {
	FullName string `json:"full_name"`
}

type customWriter struct{}

func (w customWriter) Write(p []byte) (n int, err error) {
	var resp GitHubResponse
	// Note: I had to add error handling here to get the code to work.
	marsherr := json.Unmarshal(p, &resp)
	if marsherr != nil {
		fmt.Println("Error:", marsherr)
		os.Exit(1)
	}

	for _, r := range resp {
		fmt.Println(r.FullName)
	}
	return len(p), nil
}

func main() {
	// It seems that the default (5 items per page) meant that the response body was
	// split into two parts, this meant that when the first part was unmarshalled
	// the JSON was invalid.
	resp, err := http.Get("https://api.github.com/users/microsoft/repos?page=15&per_page=4")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	writer := customWriter{}
	io.Copy(writer, resp.Body)
}
