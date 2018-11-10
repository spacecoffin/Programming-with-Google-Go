package main

import (
	"fmt"
)

/*
Write a program which prompts the user to enter a floating point
number and prints the integer which is a truncated version of the
floating point number that was entered. Truncation is the process of
removing the digits to the right of the decimal place.
*/
func main() {
	var userInput float64
	fmt.Println("Please enter a floating point number")
	_, err := fmt.Scan(&userInput)
	if err != nil {
		fmt.Println("Invalid input")
		fmt.Println(err)
		return
	}
	castAsInt := int(userInput)
	fmt.Printf("%d", castAsInt)
}
