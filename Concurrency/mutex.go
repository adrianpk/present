package main

import (
	"fmt"
	"sync"
)

// START OMIT
var counter int
var mutex sync.Mutex // HL

func increment() {
	mutex.Lock()         // Acquire the lock // HL
	defer mutex.Unlock() // Ensure the lock is released even if panic occurs // HL
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
	wg.Wait()

	fmt.Println("Counter:", counter)
}

// END OMIT
