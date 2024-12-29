package main

import (
	"fmt"
	"sync"
	"time"
)

var m sync.Mutex

var n = 10

func p() {
	m.Lock()

	fmt.Println("LOCK")
	fmt.Println(n)
	time.Sleep(1 * time.Second)

	m.Unlock()
	fmt.Println("UNLOCK")
}

func Mutex() {
	fmt.Println("FIRST")
	go p()

	fmt.Println("SECOND")
	p()

	fmt.Println("THIRD")
	time.Sleep(3 * time.Second)

	fmt.Println("DONE")
}

// The Mutex type from the sync package is used to provide a locking mechanism
// to ensure that only one goroutine can access a shared resource at a time.

// 1. Prints `FIRST`
// 2. Launches a Goroutine for p()
// 3. Prints `SECOND` and Calls p() in Main Thread
// 4. Executes p() in Main Thread
// 5. The goroutine for p() is blocked at m.Lock() because the main thread holds the mutex.
// 6. Prints `THIRD` and Sleeps for 3 seconds
// 7. After the main thread releases the lock,
// the goroutine acquires it and executes the critical section.
// 8. Prints `DONE`
