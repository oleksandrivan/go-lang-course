package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {

	var input string
	slice := make([]int, 0, 3)

	for {
		fmt.Println("Enter a number or 'x' to quit: ")
		num, err := fmt.Scan(&input)

		if num > 0 && err == nil {
			if strings.EqualFold(input, "x") {
				fmt.Println("Exiting program")
				break
			} else {
				input_int, convErr := strconv.Atoi(input)
				// No errors on conversion
				if convErr == nil {
					slice = append(slice, input_int)
					sort.Ints(slice)
					fmt.Printf("Sorted slice is %v\n", slice)
				}
			}
		}
	}
}
