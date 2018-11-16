package main

import (
	"fmt"
	"strconv"
	//"math"
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
	fmt.Println("Please enter a floating point number")
	_, err := fmt.Scanln(&userInput)
	if err != nil {
		fmt.Println("Invalid input")
		fmt.Println(err)
		return
	}
	splitInput := strings.SplitN(userInput, ".", 1)
	fmt.Printf("%v", splitInput)
	castAsInt, err := strconv.Atoi(splitInput[0])
	if err != nil {
		fmt.Println("Invalid input")
		fmt.Println(err)
		return
	}

	/*floored := math.Floor(userInput)
	castAsInt := int(floored)*/
	fmt.Printf("%d\n", castAsInt)
}
