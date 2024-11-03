package main

import (
	"fmt"
	"sync"
)

/*
The race condition happens because the counter++ operation is not atomic (indivisible) and involves reading, incrementing, and writing, 
which can be interleaved between the two goroutines. When two goroutines perform this sequence of steps on the same variable simultaneously, 
they can overwrite each other’s changes, leading to an incorrect final value of counter.
*/

func changeCounter(counter *int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 1000; i++ {
		*counter++
	}
}

func main() {
	var counter int
	var wg sync.WaitGroup

	for differentCount := 0; differentCount < 10; {
		counter = 0
		wg.Add(2)
		go changeCounter(&counter, &wg)
		go changeCounter(&counter, &wg)
		wg.Wait()
		if counter != 2000 {
			fmt.Println("Counter value:", counter)
			differentCount++
		}
	}
}
