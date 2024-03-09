package main

import "fmt"

func main() {
	ch := make(<-chan int) // Create a receive-only channel // HL

	go func() {
		// Attempting to send data to a receive-only channel will cause a compile-time error
		fmt.Println(<-ch) // Receiving data from the channel
	}()

	fmt.Println("Data received successfully")
}
