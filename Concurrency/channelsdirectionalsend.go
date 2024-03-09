package main

import "fmt"

func main() {
	ch := make(chan<- int) // Create a send-only channel

	go func() {
		ch <- 42 // Attempting to receive from a send-only channel will cause a compile-time error
	}()

	fmt.Println("Data sent successfully")
}
