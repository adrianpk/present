package main

import "fmt"

func main() {
	// Create a buffered channel
	ch := make(chan int)

	go func() {
		ch <- 42 // Send data to it
	}()

	data := <-ch // Receive data from it

	fmt.Println(data)
}
