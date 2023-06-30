package main

import (
	"fmt"
	"strconv"
)

func fizzbuzz(value int) string {
	isMatch := func(mod int) bool {
		return value%mod == 0
	}

	div3 := isMatch(3)
	div5 := isMatch(5)

	switch {
	case div3 && div5:
		return "FizzBuzz"
	case div3:
		return "Fizz"
	case div5:
		return "Buzz"
	}

	return strconv.Itoa(value)
}

func main() {
	for i := 1; i <= 100; i++ {
		fmt.Println(fizzbuzz(i))
	}
}
