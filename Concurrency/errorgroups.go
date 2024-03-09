package main

import (
	"errors"
	"fmt"
	"sync"
)

// START OMIT
func doTask(id int, eg *sync.ErrGroup) error {
	defer eg.Done() // Decrements the ErrorGroup counter when the task completes //

	if id%2 == 0 { // Simulate task error.
		return errors.New(fmt.Sprintf("Error from task %d", id))
	}

	return nil
}

func main() {
	const numTasks = 5

	eg := &sync.ErrGroup{} // Create an ErrorGroup // HL

	for i := 0; i < numTasks; i++ {
		eg.Add(1)        // Increment the ErrorGroup counter for each task // HL
		go doTask(i, eg) // Start a new goroutine for each task
	}

	// Wait for all tasks to finish and collect any errors
	err := eg.Wait() // Block until the ErrorGroup counter goes back to zero // HL
	if err != nil {
		fmt.Println("Errors occurred during tasks:")
		fmt.Println(err)
	} else {
		fmt.Println("All tasks completed successfully")
	}
}

// END OMIT
