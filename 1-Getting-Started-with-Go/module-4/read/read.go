package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Name struct
type Name struct {
	fname string
	lname string
}

/*
Write a program which reads information from a file and represents it in a slice of structs. Assume that there is a text file which contains a series of names. Each line of the text file has a first name and a last name, in that order, separated by a single space on the line.

Your program will define a name struct which has two fields, fname for the first name, and lname for the last name. Each field will be a string of size 20 (characters).

Your program should prompt the user for the name of the text file. Your program will successively read each line of the text file and create a struct which contains the first and last names found in the file. Each struct created will be added to a slice, and after all lines have been read from the file, your program will have a slice containing one struct for each line in the file. After reading all lines from the file, your program should iterate through your slice of structs and print the first and last names found in each struct.
*/
func main() {
	var nameList []Name // slice of name structs
	var fileName string
	fmt.Printf("Enter name of text file: ")
	if _, err := fmt.Scan(&fileName); err != nil {
		fmt.Println(err)
		return
	}
	f, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	scanner := bufio.NewScanner(bufio.NewReader(f))
	for scanner.Scan() {
		line := scanner.Text()
		splitLine := strings.SplitN(line, " ", 2)
		lineName := Name{fname: splitLine[0], lname: splitLine[1]}
		nameList = append(nameList, lineName)
	}
	for _, name := range nameList {
		fmt.Printf("First: %s\nLast:  %s\n\n", name.fname, name.lname)
	}
}