package main

import (
	"fmt"
	"sync"
)

// START OMIT
var counter int

func increment() {
	counter++
}

func main() {
	var wg sync.WaitGroup // We'll come back to this in the next section
	for i := 0; i < 10; i++ {
		wg.Add(1) // HL
		go func() {
			defer wg.Done() // HL
			increment()
		}()
	}

	// Wait for all goroutines to finish
	wg.Wait() // We'll come back to this in the next section

	fmt.Println("Counter:", counter)
}

// END OMIT
