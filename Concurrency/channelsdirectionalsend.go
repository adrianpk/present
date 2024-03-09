package main

import "fmt"

// START OMIT
// Function createChannel returns a send-only channel.
func createChannel() chan<- int {
	c := make(chan int) // HL

	go func() {
		for i := 1; i <= 5; i++ {
			c <- i // Send values on the channel // HL
		}
		close(c) // Close the channel after sending all values // HL
	}()

	// Returning the channel as send-only.
	return c
}

func main() {
	// Receive a send-only channel from the function.
	ch := createChannel()

	// Attempting to receive data from the send-only channel would result in a compilation error.
	// val := <-ch // HL

	// Iterate over the values sent on the send-only channel and print them.
	for val := range ch { // HL
		fmt.Println("Received:", val)
	}
}

// END OMIT
