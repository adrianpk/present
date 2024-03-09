package main

import (
	"fmt"
	"sync"
	"time"
)

// START OMIT
func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done() // Signal to the WaitGroup that this worker has finished // HL
	fmt.Printf("Worker %d starting\n", id)
	time.Sleep(time.Second) // Simulate some work
	fmt.Printf("Worker %d done\n", id)
}

func main() {
	const numWorkers = 3
	var wg sync.WaitGroup // Create a WaitGroup // HL

	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)         // Increment the WaitGroup counter for each worker // HL
		go worker(i, &wg) // Start each worker goroutine // HL
	}

	wg.Wait() // Wait for all worker goroutines to finish // HL

	fmt.Println("All workers have completed their tasks")
}

// END OMIT
