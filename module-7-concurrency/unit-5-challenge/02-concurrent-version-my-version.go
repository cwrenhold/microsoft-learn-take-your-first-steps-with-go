package main

import (
	"fmt"
	"math/rand"
	"time"
)

func fib(number float64) float64 {
	x, y := 1.0, 1.0
	for i := 0; i < int(number); i++ {
		x, y = y, x+y
	}

	r := rand.Intn(3)
	time.Sleep(time.Duration(r) * time.Second)

	return x
}

func main() {
	start := time.Now()

	size := 15
	ch := make(chan string, size)

	for i := 1; i < size; i++ {
		val := i
		go func() {
			n := fib(float64(val))
			ch <- fmt.Sprintf("Fib(%v): %v\n", val, n)
		}()
	}

	output := make([]string, size)
	for i := 1; i < size; i++ {
		output[i] = <-ch
	}

	// Sort the output - this is not necessary for the challenge.
	// It is also an alpha sort, so Fib(10) will come before Fib(2).
	// But it should look a bit nicer to read.
	for i := 1; i < size; i++ {
		for j := i + 1; j < size; j++ {
			if output[i] > output[j] {
				output[i], output[j] = output[j], output[i]
			}
		}
	}

	for i := 1; i < size; i++ {
		fmt.Printf(output[i])
	}

	elapsed := time.Since(start)
	fmt.Printf("Done! It took %v seconds!\n", elapsed.Seconds())
}
