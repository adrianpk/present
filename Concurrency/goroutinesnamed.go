package main

import (
	"fmt"
	"time"
)

// START OMIT
func main() {
	go SomeFunction("Hello from named") // HL

	// Keep the main routine running for a while to allow the goroutines to execute
	// This is a bad practice, but it's useful for this example
	// We will see a better way to handle this in the following examples
	time.Sleep(100 * time.Millisecond) // HL
}

func SomeFunction(msg string) {
	fmt.Println("Named function with parameters:", msg)
}

// END OMIT
