package main

import (
	"fmt"
	"sync"
	"time"
)

// START OMIT
func worker(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		fmt.Printf("Worker %d processing job %d\n", id, job)
		time.Sleep(time.Second) // Simulate some work
		results <- job * 2      // Send the result back to the results channel
	}
}

func main() {
	const numJobs = 5
	const numWorkers = 3

	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	var wg sync.WaitGroup
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go func(workerID int) { // Start a worker goroutine
			defer wg.Done()                 // Signal to the WaitGroup that this worker has finished
			worker(workerID, jobs, results) // Start the worker
		}(i)
	}

	// Send jobs to the worker pool
	for i := 1; i <= numJobs; i++ {
		jobs <- i // Send the job to the jobs channel
	}
	close(jobs) // Close the jobs channel to signal that all jobs have been submitted

	// Collect results from the worker pool
	go func() {
		wg.Wait()      // Wait for all workers to finish
		close(results) // Close the results channel after all workers have completed
	}()

	// Print the results
	for result := range results {
		fmt.Println("Result:", result)
	}
}

// END OMIT
