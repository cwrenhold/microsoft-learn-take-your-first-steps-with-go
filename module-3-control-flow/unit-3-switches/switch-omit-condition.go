package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	r := rand.Float64()
	switch {
	case r > 0.1:
		fmt.Println("Common case, 90% of the time")
	default:
		fmt.Println("10% of the time")
	}
}
