package main

import "fmt"

func getResponse(value int) string {
	switch {
	case value == 0:
		return "0 is neither negative nor positive"
	case value < 0:
		panic("You entered a negative number!")
	default:
		return fmt.Sprintf("You entered: %v", value)
	}
}

func main() {
	for {
		value := 0
		fmt.Print("Enter number: ")
		fmt.Scanf("%d", &value)

		output := getResponse(value)

		fmt.Println(output)
	}
}
