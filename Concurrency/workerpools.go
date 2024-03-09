package main

import (
	"fmt"
	"sync"
	"time"
)

// START OMIT
func worker(id int, jobs <-chan int, results chan<- int) { // HL
	for job := range jobs { // HL
		fmt.Printf("Worker %d processing job %d\n", id, job)
		time.Sleep(time.Second) // Simulate some work
		results <- job * 2      // Send the result back to the results channel
	} // HL
} // HL

func main() {
	const numJobs = 5
	const numWorkers = 3

	jobs := make(chan int, numJobs)    // HL
	results := make(chan int, numJobs) // HL

	var wg sync.WaitGroup              // HL
	for i := 1; i <= numWorkers; i++ { // HL
		wg.Add(1)               // HL
		go func(workerID int) { // Start a worker goroutine // HL
			defer wg.Done()                 // Signal to the WaitGroup that this worker has finished // HL
			worker(workerID, jobs, results) // Start the worker // HL
		}(i) // HL
	} // HL

	// Send jobs to the worker pool
	for i := 1; i <= numJobs; i++ { // HL
		jobs <- i // Send the job to the jobs channel // HL
	} // HL
	close(jobs) // Close the jobs channel to signal that all jobs have been submitted // HL

	// Collect results from the worker pool
	go func() { // HL
		wg.Wait()      // Wait for all workers to finish // HL
		close(results) // Close the results channel after all workers have completed // HL
	}() // HL

	// Print the results
	for result := range results { // HL
		fmt.Println("Result:", result) // HL
	} // HL
}

// END OMIT
