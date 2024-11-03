package main

import (
	"fmt"
	"strings"
	"bufio"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter the string for finding I A N: ")

	input, err := reader.ReadString('\n')
	input = strings.Trim(input, "\n")

	if err == nil {
		fmt.Printf("The entered string is '%s' \n", input)

		if findIan(&input) {
			fmt.Println("Found!")
		} else {
			fmt.Println("Not Found!")
		}

	} else {
		fmt.Println("Error found!", err)
	}
}

func findIan(s *string) bool {

	var firstI, containsA, endN bool

	for index, value := range *s {
		switch {

		case index == 0:
			if strings.EqualFold(string(value), "I") {
				firstI = true
				continue
			} else {
				return false
			}

		case index < len(*s)-1:
			if strings.EqualFold(string(value), "A") {
				containsA = true
				continue
			}

		case index == len(*s)-1:
			if strings.EqualFold(string(value), "N") {
				endN = true
				continue
			} else {
				return false
			}

		}
	}
	return firstI && containsA && endN
}
