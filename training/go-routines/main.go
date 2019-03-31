package main

import (
	"fmt"
	"sync"
	"time"
)

var smallDelay int64 = 200
var largeDelay int64 = 500

// We will use WaitGroup to ensure that all related
// go-routines will be finished before main() go routine
// exits. WaitGroup.Add() should be always the first step.
// We register here how many child go-routines we expect to be
// done. WaitGroup.Wait() waits until those routines will
// be completed. It is responsibility of the programmer to
// explicitly tell when those 'jobs' will be finished.
// We can achieve this with WaitGroup.Done()
var wg sync.WaitGroup

func main() {
	wg.Add(2)
	go smallDelayCounter()
	go largeDelayCounter()
	wg.Wait()
}

func smallDelayCounter() {
	for i := 0; i <= 50; i++ {
		fmt.Println("Small Delay: ", i)
		time.Sleep(time.Duration(smallDelay) * time.Millisecond)
	}
	wg.Done()
}

func largeDelayCounter() {
	for i := 0; i <= 20; i++ {
		fmt.Println("Large Delay: ", i)
		time.Sleep(time.Duration(largeDelay) * time.Millisecond)
	}
	wg.Done()
}
