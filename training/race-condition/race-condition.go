package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Race conditions occur when two or more Goroutines try to
// use a piece of shared data at the same time.
// One way we could prevent a data race is to ensure that if
// one Goroutine is using 'counter' variable, then all other
// Goroutines are prevented (or mutually excluded) from using
// it at the same time. This is where Mutex structure starts to
// shine. One of the technique for preventing data races is
// optimistic locking. While one Goroutine holds the lock,
// all other Goroutines are not allowed to execute any lines of
// code protected by the same mutex. They are forced to wait
// until the lock is yielded.

// Counter is a variable declared in the global scope.
// It means it will be accessed by both running go routines
// at the same time. This is where race condition will happen.
// Go routines are like lightweight independent threads, so
// they cannot see global counter mutations on real time.
// They overwrite each other all the time.

var counter int
var wg sync.WaitGroup

func main() {
	wg.Add(2)
	go incrementCounter()
	go incrementCounter()
	wg.Wait()
}

func incrementCounter() {
	for i := 0; i < 20; i++ {
		x := counter
		x++
		time.Sleep(time.Duration(rand.Intn(3)) * time.Millisecond)
		counter = x
		fmt.Println("Counter is: ", counter)
	}
	wg.Done()
}
