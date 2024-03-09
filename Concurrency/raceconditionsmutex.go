package main

import (
	"fmt"
	"sync"
	"time"
)

// START OMIT
func main() {
	var counter int
	var mutex sync.Mutex // Create a Mutex

	for i := 0; i < 1000; i++ {
		go func() {
			mutex.Lock() // Acquire the lock before accessing the counter // HL
			counter++
			mutex.Unlock() // Release the lock after accessing the counter // HL
		}()
	}

	time.Sleep(time.Second) // Allowing time for goroutines to finish

	fmt.Println("Counter:", counter)
}

// END OMIT
