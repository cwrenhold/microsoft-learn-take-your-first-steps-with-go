package main

import (
	"fmt"
)

func main() {
	var value string

	fmt.Print("Enter a Roman numeral to convert: ")
	fmt.Scanf("%s", &value)

	testValues := []string{value, "MCLX", "MCMXCIX", "MCMZ"}

	for _, testValue := range testValues {
		mine := mySolution(testValue)
		theirs := convertRomanToInteger(testValue)
		fmt.Printf("%s is %d on my solution, or %d with Microsoft's\n", testValue, mine, theirs)
	}
}

func mySolution(numeral string) int {
	romanMap := map[rune]int{
		'M': 1000,
		'D': 500,
		'C': 100,
		'L': 50,
		'X': 10,
		'V': 5,
		'I': 1,
	}

	total := 0

	// Set the lastValue to the largest value in the map.
	lastValue := 0
	for _, value := range romanMap {
		if value > lastValue {
			lastValue = value
		}
	}

	for _, char := range numeral {
		value, exists := romanMap[char]

		if !exists {
			fmt.Printf("Error: The roman numeral %s has an invalid character: %c\n", numeral, char)
			return 0
		}

		total += value

		if value > lastValue {
			// We've already added the last value, so we need to subtract it twice.
			total -= lastValue * 2
		}

		lastValue = value
	}

	return total
}

func convertRomanToInteger(numeral string) int {
	romanMap := map[rune]int{
		'M': 1000,
		'D': 500,
		'C': 100,
		'L': 50,
		'X': 10,
		'V': 5,
		'I': 1,
	}

	// Create a slice to store the values of the roman numerals.
	// The slice will be one larger than the length of the numeral
	// to allow for the last value to be zero.
	arabicVals := make([]int, len(numeral)+1)

	// Iterate over the numeral and convert each digit to an integer.
	for index, digit := range numeral {
		// Check that the digit is valid. If it is, add it to the slice.
		if val, present := romanMap[digit]; present {
			arabicVals[index] = val
		} else {
			fmt.Printf("Error: The roman numeral %s has a bad digit: %c\n", numeral, digit)
			return 0
		}
	}

	// Initialise the total to zero.
	total := 0

	// Iterate over the found values and add them to the total.
	for index := 0; index < len(numeral); index++ {
		// If the value is less than the next value, then it is a negative value.
		// E.g. IV = 5 - 1 = 4
		if arabicVals[index] < arabicVals[index+1] {
			arabicVals[index] = -arabicVals[index]
		}
		total += arabicVals[index]
	}

	return total
}
