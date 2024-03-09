package main

import "fmt"

// START OMIT
func main() {
	ch := make(chan int) // Create a bidirectional channel // HL

	go func() {
		ch <- 42 // Send data to the channel // HL
	}()

	data := <-ch // Receive data from the channel // HL
	fmt.Println(data)
}

// END OMIT
