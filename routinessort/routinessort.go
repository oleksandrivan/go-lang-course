package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"
	"sort"
)

func main() {
	var wg sync.WaitGroup

	inputSlice := UserInput()
	chunks := SplitIntoChunks(inputSlice, 4)
	wg.Add(4)
	for i, s := range chunks {
		fmt.Printf("Slice to sort %d: %v\n", i, s)
		go sortSlice(&s, &wg)
	}
	wg.Wait()
	for i, s := range chunks {
		fmt.Printf("Sorted slice %d: %v\n", i, s)
	}
	sorted := mergeAll(chunks)
	fmt.Println("Sorted array:", sorted)
}

func merge(left, right []int) []int {
	result := make([]int, 0, len(left)+len(right))
	i, j := 0, 0
	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}
	result = append(result, left[i:]...)
	result = append(result, right[j:]...)
	return result
}

func mergeAll(parts [][]int) []int {
	merged := merge(parts[0], parts[1])
	merged = merge(merged, parts[2])
	merged = merge(merged, parts[3])
	return merged
}

func sortSlice(s *[]int, wg *sync.WaitGroup) {
	defer wg.Done()
	sort.Ints(*s)
}

func SplitIntoChunks(inputSlice []int, numberOfChunks int) [][]int {
	var chunks [][]int
	chunkSize := (len(inputSlice) + numberOfChunks - 1) / numberOfChunks // ceiling division
	for i := 0; i < len(inputSlice); i += chunkSize {
		end := i + chunkSize
		if end > len(inputSlice) {
			end = len(inputSlice)
		}
		chunks = append(chunks, inputSlice[i:end])
	}
	return chunks
}

func UserInput() []int {
	inputSlice := make([]int, 0, 10)
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Input space separated integers to sort:")
	inputString, err := reader.ReadString('\n')
	fields := strings.Fields(inputString)
	if err != nil {
		log.Fatalln("Error at reading input")
	}
	for _, field := range fields {
		num, err := strconv.Atoi(field)
		if err != nil {
			log.Fatalf("Error converting input to integer: %v", err)
		}
		inputSlice = append(inputSlice, num)
	}
	return inputSlice
}
