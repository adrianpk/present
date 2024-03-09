package main

import (
	"fmt"
)

// START OMIT
// Function sendData sends values on a bidirectional channel.
func sendData(ch chan<- int) {
	for i := 1; i <= 5; i++ {
		ch <- i // Send values on the channel // HL
	}
	close(ch) // Close the channel after sending all values
}

// Function receiveData receives values from a bidirectional channel.
func receiveData(ch <-chan int) {
	for val := range ch { // Get values from the channel // HL
		fmt.Println("Received:", val)
	}
}

func main() {
	// Create a bidirectional channel.
	ch := make(chan int)

	// Start goroutines to send and receive data on the channel.
	go sendData(ch) // HL
	receiveData(ch) // HL
}

// END OMIT
