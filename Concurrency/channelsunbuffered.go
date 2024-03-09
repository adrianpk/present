package main

import "fmt"

// START OMIT
func main() {
	ch := make(chan int, 10) // Create an unbuffered channel with buffer of 10 // HL

	for i := 0; i < 10; i++ {
		go func(i int) {
			ch <- i // Send data to the channel asynchronously // HL
		}(i)
	}

	for i := 0; i < 10; i++ {
		data := <-ch // Receive data from the channel // HL
		fmt.Println(data)
	}
}

// END OMIT
