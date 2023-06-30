package main

import "fmt"

func main() {
	var a [3]int
	a[1] = 10
	a[2] = 20
	fmt.Println(a[0])
	fmt.Println(a[1])
	fmt.Println(a[len(a)-1])
}
