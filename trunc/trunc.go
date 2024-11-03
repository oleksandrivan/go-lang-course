package main

import "fmt"

func main() {
	var input float32


	fmt.Println("Enter a floating point number: ")
	num, error := fmt.Scan(&input)

	if num > 0 && error == nil {
		fmt.Println("Truncated number is: ", int(input))
	} else {
		fmt.Println("Invalid input")
	}
}
