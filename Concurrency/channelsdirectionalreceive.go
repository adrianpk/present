package main

import "fmt"

// START OMIT
// Function F returns a read-only channel.
func someFunction() <-chan int { // HL
	c := make(chan int)

	go func() {
		defer close(c)
		c <- 123 // Send a value on the channel // HL
	}()

	// Returning the channel as read-only.
	return c
}

func main() {
	// Receive a read-only channel from function F.
	ch := someFunction()

	// Attempting to send data on the receive-only channel would result in a compile-time error.
	// ch <- 456 // HL

	val := <-ch // Receive the value from the channel // HL
	fmt.Println("Received:", val)
}

// END OMIT
