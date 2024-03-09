package main

import (
	"fmt"
)

// START OMIT
func main() {
	const numGoroutines = 1000

	counter := make(chan int) // A channel for incrementing the counter
	done := make(chan bool)   // A channel to signal when all goroutines are done

	for i := 0; i < numGoroutines; i++ { // Start numGoroutines goroutines
		go func() { // Start each goroutine
			counter <- 1 // Increment the counter
			done <- true // Signal that this goroutine is done
		}()
	}

	// Create a separate goroutine to close the done channel
	go func() {
		for i := 0; i < numGoroutines; i++ {
			<-done // Receive from the done channel
		}
		close(done) // Close the done channel after all signals are received
	}()

	// Calculate the final value of the counter
	var finalCounter int
	for i := 0; i < numGoroutines; i++ {
		finalCounter += <-counter // Receive from the counter channel and accumulate
	}

	close(counter) // Close the counter channel

	fmt.Println("Counter:", finalCounter)
}

// END OMIT
