// Select

package main

import (
	"fmt"
	"time"
)

// START OMIT
func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() { // Send "ping" message every 500 milliseconds
		for {
			time.Sleep(500 * time.Millisecond)
			ch1 <- "ping"
		}
	}()

	go func() { // Send "pong" message every 1 second
		for {
			time.Sleep(1 * time.Second)
			ch2 <- "pong"
		}
	}()

	for {
		select { // Receive messages from ch1 and ch2 // HL
		case msg1 := <-ch1: // HL
			fmt.Println("Received message from ch1:", msg1)
		case msg2 := <-ch2: // HL
			fmt.Println("Received message from ch2:", msg2)
		}
	}
}

// END OMIT
