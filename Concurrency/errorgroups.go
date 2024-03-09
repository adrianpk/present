package main

import (
	"errors"
	"fmt"
	"os"

	"golang.org/x/sync/errgroup"
)

// START OMIT
func doTask(id int, eg *errgroup.Group) error {
	if id%2 == 0 { // Simulate task error.
		return errors.New(fmt.Sprintf("Error from task %d", id))
	}
	return nil
}

func main() {
	const numTasks = 5

	eg := &errgroup.Group{} // Create an ErrorGroup. // HL

	for i := 0; i < numTasks; i++ {
		id := i              // Capture the loop variable
		eg.Go(func() error { // HL
			return doTask(id, eg)
		})
	}

	// Wait for all tasks to finish and collect any errors
	err := eg.Wait() // Block until all tasks complete // HL
	if err != nil {
		fmt.Println("Errors occurred during tasks:")
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("All tasks completed successfully")
}

// END OMIT
