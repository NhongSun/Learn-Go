package main

import (
	"fmt"
	"sync"
)

func Once() {
	var once sync.Once
	var wg sync.WaitGroup

	initialize := func() {
		fmt.Println("Initializing only once")
	}

	doWork := func(workerId int) {
		defer wg.Done()
		fmt.Printf("Worker %d started\n", workerId)
		once.Do(initialize) // This will only be executed once
		fmt.Printf("Worker %d done\n", workerId)
	}

	numWorkers := 5
	wg.Add(numWorkers)

	// Launch several goroutines
	for i := 0; i < numWorkers; i++ {
		go doWork(i)
	}

	// Wait for all goroutines to complete
	wg.Wait()
	fmt.Println("All workers completed")
}

// No matter how many times once.Do is called,
// the function passed to once.Do will only be executed once.
// the other time it is called, it will be ignored and the program will continue.
