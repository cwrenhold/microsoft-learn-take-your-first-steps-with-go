package main

import (
	"fmt"
	"time"
)

func main() {
	switch time.Now().Weekday().String() {
	case "Monday", "Tuesday", "Wednesday", "Thursday", "Friday":
		fmt.Println("It's time to learn some Go.")
	default:
		fmt.Println("It's the weekend, time to rest!")
	}

	fmt.Println(time.Now().Weekday().String())
}
