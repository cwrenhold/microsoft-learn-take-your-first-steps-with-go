package main

import "fmt"

func main() {
	// Note: The slice is declare with the [] operator and is of length 12 as
	//       we have supplied 12 values. The capacity is also 12 to match.
	months := []string{"January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"}
	fmt.Println(months)
	fmt.Println("Length:", len(months))
	fmt.Println("Capacity:", cap(months))
}
