package main

import (
	"fmt"
	"sync"
)

// Counter struct holds a value and a mutex
type Counter struct {
	value int
	mu    sync.Mutex
}

// Increment method increments the counter's value safely using the mutex
func (c *Counter) Increment() {
	c.mu.Lock()   // Lock the mutex before accessing the value
	c.value++     // Increment the value
	c.mu.Unlock() // Unlock the mutex after accessing the value
}

// Value method returns the current value of the counter
func (c *Counter) Value() int {
	return c.value
}

func Mutex2() {
	var wg sync.WaitGroup
	counter := Counter{}

	// Start 10 goroutines
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 100; j++ {
				counter.Increment()
			}
		}()
	}

	wg.Wait() // Wait for all goroutines to finish
	fmt.Println("Final counter value:", counter.Value())
}

// Since the counter is protected by a mutex, the increments are done safely.
// The final value is 1000 because 10 goroutines each increment the counter by 100.
// But if the counter was not protected by a mutex, the final value MIGHT be less than 1000,
