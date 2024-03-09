package main

import (
	"fmt"
	"time"
)

// START OMIT
func main() {
	var counter int

	for i := 0; i < 1000; i++ {
		go func() {
			counter++ // Counter is incremented by multiple goroutines whitout synchronization
		}()
	}

	time.Sleep(time.Second) // Allowing time for goroutines to finish

	fmt.Println("Counter:", counter)
}

// END OMIT
