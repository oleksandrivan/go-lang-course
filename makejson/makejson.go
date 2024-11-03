package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	personMap := make(map[string]string)

	fmt.Println("Enter a name: ")
	inputName, errName := reader.ReadString('\n')
	inputName = strings.Trim(inputName, "\n")

	if errName == nil {
		personMap["name"] = inputName
	}

	fmt.Println("Enter an address: ")
	inputAddress, errAddress := reader.ReadString('\n')
	inputAddress = strings.Trim(inputAddress, "\n")

	if errAddress == nil {
		personMap["address"] = inputAddress
	}

	bArray, jsonErr := json.Marshal(personMap)
	if jsonErr == nil {
		fmt.Printf("%s\n", bArray)
	}
}
