package main

import (
	"fmt"
	"os"
	"strings"
)

type Name struct {
	fname string
	lname string
}

func main() {

	var inputFile string

	_, scanErr := fmt.Scan(&inputFile)

	if scanErr != nil {
		return
	}

	fileByteArray, fileErr := os.ReadFile(inputFile)
	if fileErr != nil {
		return
	}
	linesArray := strings.Split(string(fileByteArray), "\n")

	nameSlice := make([]Name, 0, 5)

	for _, line := range linesArray {
		nameSplit := strings.Split(line, " ")
		nameSlice = append(nameSlice, Name{fname: nameSplit[0], lname: nameSplit[1]})
	}

	for _, name := range nameSlice {
		fmt.Printf("First name '%s' Last Name '%s'\n", name.fname, name.lname)
	}
}
