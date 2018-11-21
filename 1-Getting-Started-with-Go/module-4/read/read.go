package main

import (
	"fmt"
	"io"
	"os"
	)

// Name struct
type Name struct {
	fname string
	lname  string
}

/*
Write a program which reads information from a file and represents it in a slice of structs. Assume that there is a text file which contains a series of names. Each line of the text file has a first name and a last name, in that order, separated by a single space on the line.

Your program will define a name struct which has two fields, fname for the first name, and lname for the last name. Each field will be a string of size 20 (characters).

Your program should prompt the user for the name of the text file. Your program will successively read each line of the text file and create a struct which contains the first and last names found in the file. Each struct created will be added to a slice, and after all lines have been read from the file, your program will have a slice containing one struct for each line in the file. After reading all lines from the file, your program should iterate through your slice of structs and print the first and last names found in each struct.
*/

const NoSpaceError = "No space found after first name."

func isSpace(r rune) bool {
	if r == ' ' {
		return true
	}
	return false
}

func isNewline(r rune) bool {
	switch r {
	case '\n', '\r':
		return true
	}
	return false
}

func findSep(barr []byte, sepComp func(r rune) bool) {
	for index, char := range(fBarr) {
		if char == space {

		}
	}
}

func readFirst(fileHandle os.File) (string, error) {
	var
	fBarr := make([]byte, 0, 21)
	nBytesRead, err := fileHandle.read(fBarr)
	if err != nil {
		return nil, err
	}


	if string(fBarr[nBytesRead]) != " " {
		return nil, NoSpaceError
	}
	fString := string(fBarr[:nBytesRead])
	return fString, nil
}

func readLast(fileHandle os.File) (string, error) {
	fBarr := make([]byte, 0, 21)
	nBytesRead, err := fileHandle.read(fBarr)
	if err != nil & err != io.EOF{
		return nil, err
	}
	if string(fBarr[nBytesRead]) != " " {
		return nil, NoSpaceError
	}
	fString := string(fBarr[:nBytesRead])
	return fString, nil
}

func readLine(fileHandle os.File) (Name, error) {
	fBarr := [20]byte

	lBarr := [20]byte

}

/*
I decided to complete this assignment with the fewest possible imports. Maybe I miss my university days of writing C.
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



}
