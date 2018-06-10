package main

import (
	"fmt"
	"time"
	"sync"
	"math/rand"
)

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

