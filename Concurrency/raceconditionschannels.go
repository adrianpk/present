package main

import (
	"fmt"
)

// START OMIT
func main() {
	const numGoroutines = 1000

	counter := make(chan int) // A channel for incrementing the counter
	done := make(chan bool)   // A channel to signal when all finished

	for i := 0; i < numGoroutines; i++ { // Start numGoroutines goroutines // HL
		go func() { // HL
			counter <- 1 // Increment the counter // HL
			done <- true // Signal that this goroutine is done // HL
		}() // HL
	} // HL

	for i := 0; i < numGoroutines; i++ { // Wait for all goroutines to finish
		<-done // Receive from the done channel // HL
	}

	close(counter) // Close the counter channel // HL

	// Calculate the final value of the counter
	var finalCounter int
	for n := range counter { // Range over the counter channel values // HL
		finalCounter += n
	}

	fmt.Println("Counter:", finalCounter)
}

// END OMIT
