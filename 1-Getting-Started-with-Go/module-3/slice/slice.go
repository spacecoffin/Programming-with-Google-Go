package main

import (
	"fmt"
	"sort"
	"strconv"
)

func printPrompt() {
	fmt.Println("Enter an integer to add it to the sorted list -or- 'X' to exit.")
}

/*
Write a program which prompts the user to enter integers and stores the integers in a sorted slice. The program should be written as a loop. Before entering the loop, the program should create an empty integer slice of size (length) 3. During each pass through the loop, the program prompts the user to enter an integer to be added to the slice. The program adds the integer to the slice, sorts the slice, and prints the contents of the slice in sorted order. The slice must grow in size to accommodate any number of integers which the user decides to enter. The program should only quit (exiting the loop) when the user enters the character ‘X’ instead of an integer.
*/
func main() {
	var userInput string
	slice := make([]int, 0, 3)

	for {
		printPrompt()
		_, err := fmt.Scan(&userInput)
		if err != nil {
			fmt.Println(err)
			break
		}
		if userInput == "X" {
			fmt.Println("Exiting.")
			break
		}
		enteredInt, err := strconv.Atoi(userInput)
		if err != nil {
			fmt.Println(err)
			continue
		}
		slice = append(slice, enteredInt)
		if !sort.SliceIsSorted(slice, func(i int, j int) bool { return slice[i] < slice[j] }) {
			sort.SliceStable(slice, func(i int, j int) bool { return slice[i] < slice[j] })
		}
		fmt.Println("Sorted slice: ", slice)
	}
}
