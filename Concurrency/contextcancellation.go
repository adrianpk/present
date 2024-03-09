package main

import (
	"context"
	"fmt"
	"time"
)

// START OMIT
func task1(ctx context.Context) {
	for {
		select { // HL
		case <-ctx.Done(): // HL
			fmt.Println("Task 1: Context cancelled")
			return
		default: // HL
			// Simulate task execution
			time.Sleep(500 * time.Millisecond)
			// Simulated error condition
			if time.Now().Unix()%2 == 0 {
				fmt.Println("Task 1: Error occurred")
				return
			}
			fmt.Println("Task 1: Processing...")
		} // HL
	}
}

func task2(ctx context.Context) {
	for {
		select {
		case <-ctx.Done(): // HL
			fmt.Println("Task 2: Context cancelled")
			return
		default: // HL
			// Simulate task execution
			time.Sleep(700 * time.Millisecond)
			fmt.Println("Task 2: Processing...")
		} // HL
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() // Cancel the context when main function exits // HL

	// Start task1 and task2 goroutines
	go task1(ctx) // HL
	go task2(ctx) // HL

	// Create a ticker to cancel the context after 3 seconds
	ticker := time.NewTicker(3 * time.Second) // HL
	defer ticker.Stop()

	select { // HL
	case <-ticker.C: // HL
		fmt.Println("Main: Timeout reached, cancelling context")
		cancel() // Cancel the context after the timeout // HL
	case <-ctx.Done(): // HL
		fmt.Println("Main: Task completed, context cancelled")
	} // HL

	// Allow some time for the goroutines to finish
	time.Sleep(1 * time.Second)
	fmt.Println("Main: Exiting...")
}

// END OMIT
