package main

import (
	"fmt"
	"strings"
)

/*
Write a program which prompts the user to enter a string. The program searches through the entered string for the characters ‘i’, ‘a’, and ‘n’. The program should print “Found!” if the entered string starts with the character ‘i’, ends with the character ‘n’, and contains the character ‘a’. The program should print “Not Found!” otherwise. The program should not be case-sensitive, so it does not matter if the characters are upper-case or lower-case.
*/

// Findian : Parses string to determine if it matches conditions.
func Findian(input string) string {
	lowerInput := strings.ToLower(input)
	startsWithI := strings.HasPrefix(lowerInput, "i")
	endsWithN := strings.HasSuffix(lowerInput, "n")
	containsA := strings.Contains(lowerInput, "a")
	if startsWithI && endsWithN && containsA {
		return "Found!"
	}
	return "Not Found!"
}

func main() {
	var userInput string
	fmt.Println("Enter a string:")
	_, err := fmt.Scan(&userInput)
	if err != nil {
		fmt.Println(err)
		return
	}
	result := Findian(userInput)
	fmt.Println(result)
}
