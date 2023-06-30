package main

import "fmt"

func main() {
	letters := []string{"A", "B", "C", "D", "E"}
	remove := 2

	if remove < len(letters) {

		fmt.Println("Before", letters, "Remove ", letters[remove])

		// This replaces the slice using a concatenation of the letters
		// before the one we want to remove and the letters after the one we
		// want to remove.
		letters = append(letters[:remove], letters[remove+1:]...)

		fmt.Println("After", letters)
	}

}
