package main

import (
	"fmt"
	"strconv"
	"strings"
)

/*
Write a program which prompts the user to enter a floating point
number and prints the integer which is a truncated version of the
floating point number that was entered. Truncation is the process of
removing the digits to the right of the decimal place.
*/
func main() {
	var userInput string
	fmt.Printf("Please enter a floating point number: ")
	if _, err := fmt.Scanln(&userInput); err != nil {
		fmt.Println(err)
		return
	}
	if _, err := strconv.ParseFloat(userInput, 64); err != nil {
		fmt.Println(err)
		return
	}
	splitInput := strings.SplitN(userInput, ".", 2)
	castAsInt, err := strconv.Atoi(splitInput[0])
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%d\n", castAsInt)
}
