package main

import "fmt"

// START OMIT
func main() {
	ch := make(chan int, 1) // Create a buffered channel with capacity 1 // HL

	ch <- 42 // Send data to the channel asynchronously // HL

	data := <-ch // Receive data from the channel // HL
	fmt.Println(data)
}

// END OMIT
