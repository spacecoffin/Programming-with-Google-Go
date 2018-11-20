package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

/*
Write a program which prompts the user to first enter a name, and then enter an address. Your program should create a map and add the name and address to the map using the keys "name" and "address", respectively. Your program should use Marshal() to create a JSON object from the map, and then your program should print the JSON object.
*/
func main() {
	personMap := make(map[string]string)
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Please enter a name: ")
	personName, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Print("Please enter an address: ")
	personAddress, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		return
	}
	personMap["name"] = strings.Trim(personName, "\n")
	personMap["address"] = strings.Trim(personAddress, "\n")
	barr, err := json.Marshal(personMap)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(barr))
}
