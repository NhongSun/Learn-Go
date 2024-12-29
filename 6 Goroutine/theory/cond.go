package main

import (
	"fmt"
	"sync"
	"time"
)

func Cond() {
	// Create a new condition variable
	var mutex sync.Mutex
	cond := sync.NewCond(&mutex)

	// A shared resource
	ready := false

	// A goroutine that waits for a condition
	go func() {
		fmt.Println("Goroutine: Waiting for the condition...")

		mutex.Lock()
		for !ready {
			cond.Wait() // Wait for the condition
		}
		fmt.Println("Goroutine: Condition met, proceeding...")
		mutex.Unlock()
	}()

	// Simulate some work (e.g., loading resources)
	time.Sleep(2 * time.Second)

	// Signal the condition
	mutex.Lock()
	ready = true
	cond.Signal() // Signal one waiting goroutine
	mutex.Unlock()
	fmt.Println("Push signal !")

	// Give some time for the goroutine to complete
	time.Sleep(1 * time.Second)
	fmt.Println("Work is done.")
}

// When goroutine starts, it prints `Goroutine: Waiting for the condition...`
// lock the mutex and check the ready flag in a loop.
// If the ready flag is false, the goroutine waits for the condition to be met using cond.Wait().
// goroutine enters a waiting state and releases the mutex lock.
// The main thread sleeps for 2 seconds.
// After sleeping, the main thread locks the mutex and sets ready to true.
// then prints `Push signal !`
// It calls cond.Signal() to wake up one waiting goroutine
// The goroutine wakes up, acquires the mutex lock, and prints `Goroutine: Condition met, proceeding...`
// The main thread sleeps for 1 second to give the goroutine time to complete.
// Finally, the main thread prints `Work is done.`
