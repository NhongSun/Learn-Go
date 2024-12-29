package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done() // Decrement the counter when the goroutine completes

	fmt.Printf("Worker %d starting\n", id)

	// Simulate some work by sleeping
	sleepDuration := time.Duration(rand.Intn(1000)) * time.Millisecond
	time.Sleep(sleepDuration)

	fmt.Printf("Worker %d done\n", id)
}

func WaitGroup() {
	var wg sync.WaitGroup

	// Launch several goroutines and increment the WaitGroup counter for each
	wg.Add(5)
	for i := 1; i <= 5; i++ {
		go worker(i, &wg)
	}

	wg.Wait() // Block until the WaitGroup counter goes back to 0; all workers are done

	fmt.Println("All workers completed")
}

// The WaitGroup type from the sync package is used to wait for a collection of goroutines to finish executing.
// The main goroutine calls Add to set the number of goroutines to wait for.
// Then each worker goroutine runs and calls Done when it completes.
// Finally, Wait blocks until the WaitGroup counter goes back to 0; all workers are done.

// When each worker goroutine begins execution, it prints `Worker X starting`
// Since the workers start concurrently, their "starting" messages may not appear in sequential order,
// depending on how the Go runtime schedules the goroutines.
// Each worker simulates work by sleeping for a random duration (0–999 ms).
// After completing the sleep, it prints `Worker X done`
// Once all workers complete their tasks and call wg.Done()
// the main thread unblocks at wg.Wait() and prints `All workers completed`
