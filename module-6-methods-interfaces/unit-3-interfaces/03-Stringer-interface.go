package main

import "fmt"

type Person struct {
	Name, Country string
}

// Note: The String() method is part of the fmt.Stringer interface.
//
//	The fmt package (and many others) look for this interface
//	and if it is found, it will use the String() method to
//	print the value of the type.
func (p Person) String() string {
	return fmt.Sprintf("%v is from %v", p.Name, p.Country)
}

func main() {
	rs := Person{"John Doe", "USA"}
	ab := Person{"Mark Collins", "United Kingdom"}
	fmt.Printf("%s\n%s\n", rs, ab)
}
