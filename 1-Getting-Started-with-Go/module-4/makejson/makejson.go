package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

// Person struct to unmarshall JSON into
type Person struct {
	Name    string
	Address string
}

/*
Write a program which prompts the user to first enter a name, and then enter an address. Your program should create a map and add the name and address to the map using the keys "name" and "address", respectively. Your program should use Marshal() to create a JSON object from the map, and then your program should print the JSON object.
*/
func main() {
	var person Person
	personMap := make(map[string]string)
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Please enter a name: ")
	personName, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Please enter an address: ")
	personAddress, err := reader.ReadString('\n')
	if err != nil {
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
	if err := json.Unmarshal(barr, &person); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("\nUnmarshalled JSON as Go object:\n%+v\n", person)
}
