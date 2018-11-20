package main

import (
	"encoding/json"
	"fmt"
)

/*
Write a program which prompts the user to first enter a name, and then enter an address. Your program should create a map and add the name and address to the map using the keys "name" and "address", respectively. Your program should use Marshal() to create a JSON object from the map, and then your program should print the JSON object.
*/
func main() {
	var personName, personAddress string
	personMap := make(map[string]string)
	fmt.Println("Please enter a name: ")
	if _, err := fmt.Scanln(&personName); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Please enter an address: ")
	if _, err := fmt.Scanln(&personAddress); err != nil {
		fmt.Println(err)
		return
	}
	personMap["name"] = personName
	personMap["address"] = personAddress
	barr, err := json.Marshal(personMap)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(barr)
}
