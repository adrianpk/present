package main

import (
	"fmt"
	"sync"
	"time"
)

// START OMIT
func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done() // Decrements the WaitGroup counter on exit // HL
	fmt.Printf("Worker %d starting\n", id)
	time.Sleep(time.Second) // Simulate some work
	fmt.Printf("Worker %d done\n", id)
}

func main() {
	var wg sync.WaitGroup // Create a WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1) // Increment the WaitGroup counter before each goroutine starts // HL
		go worker(i, &wg)
	}

	wg.Wait() // Block until the WaitGroup counter goes back to zero // HL

	fmt.Println("All workers have finished their tasks")
}

// END OMIT
