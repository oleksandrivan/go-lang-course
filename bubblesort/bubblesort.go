package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

/*
Write a Bubble Sort program in Go. The program
should prompt the user to type in a sequence of up to 10 integers. The program
should print the integers out on one line, in sorted order, from least to
greatest. Use your favorite search tool to find a description of how the bubble
sort algorithm works.

As part of this program, you should write a
function called BubbleSort() which
takes a slice of integers as an argument and returns nothing. The BubbleSort() function should modify the slice so that the elements are in sorted
order.

A recurring operation in the bubble sort algorithm is
the Swap operation which swaps the position of two adjacent elements in the
slice. You should write a Swap() function which performs this operation. Your Swap()
function should take two arguments, a slice of integers and an index value i which
indicates a position in the slice. The Swap() function should return nothing, but it should swap
the contents of the slice in position i with the contents in position i+1.

Submit your Go program source code.
*/

func main() {
	mySlice := UserInput()
	BubbleSort(mySlice)
	fmt.Println(mySlice)
}

func UserInput() []int {
	inputSlice := make([]int, 0, 10)
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Input up to 10 space separated integers to sort:")
	inputString, err := reader.ReadString('\n')
	fields := strings.Fields(inputString)
	if err != nil {
		log.Fatalln("Error at reading input")
	}
	for i := 0; i < len(fields) && i < 10; i++ {
		num, err := strconv.Atoi(fields[i])
		if err != nil {
			log.Fatalf("Error converting input to integer: %v", err)
		}
		inputSlice = append(inputSlice, num)
	}
	return inputSlice
}

func BubbleSort(slice []int) {
	for i := 0; i < len(slice)-1; i++ {
		for j := 0; j < len(slice)-1-i; j++ {
			if slice[j] > slice[j+1] {
				Swap(slice, j)
			}
		}
	}
}

func Swap(slice []int, indexToSwap int) {
	temp := slice[indexToSwap]
	slice[indexToSwap] = slice[indexToSwap+1]
	slice[indexToSwap+1] = temp
}
