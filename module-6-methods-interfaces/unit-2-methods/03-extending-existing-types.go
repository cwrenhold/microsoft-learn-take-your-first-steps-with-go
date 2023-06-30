package main

import (
	"fmt"
	"strings"
)

// Note: This is type acts as a wrapper for the string type.
// We cannot add methods to the string type directly as it
// isn't local, so we create a new type and add methods to that.
type upperstring string

func (s upperstring) Upper() string {
	return strings.ToUpper(string(s))
}

func main() {
	s := upperstring("Learning Go!")
	fmt.Println(s)
	fmt.Println(s.Upper())
}
