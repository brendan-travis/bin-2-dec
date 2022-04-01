package main

import (
	"fmt"
	"regexp"
)

func main() {
	fmt.Println("BIN 2 DEC")

	var input string

	fmt.Println("\nEnter a decimal (up to 8 digits):")
	fmt.Scanf("%s", &input)
	fmt.Println("You have entered ", input)

	valid, _ := regexp.MatchString("[10]+", input)
	if !valid {
		fmt.Println("\nYou can only enter 1s and 0s.")
		return
	}

	if len(input) > 8 {
		fmt.Println("\nThe maximum length supported is 8 digits.")
		return
	}
}
