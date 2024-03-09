package main

import (
	"fmt"
	"time"
)

// START OMIT
func main() {
	// Anonymous function without parameters
	// This not blocks the main routine
	go func() { // HL
		fmt.Println("Without parameters")
	}()

	// Anonymous function with parameters
	// This not blocks the main routine
	go func(msg string) { // HL
		fmt.Println("With parameters:", msg) // HL
	}("Hello from anonymous") // HL

	// Keep the main routine running for a while to allow the goroutines to execute
	// This is a bad practice, but it's useful for this example
	// We will see a better way to handle this in the following examples
	time.Sleep(100 * time.Millisecond) // HL
}

// END OMIT
