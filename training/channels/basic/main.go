package main

import (
	"fmt"
	"time"
)

// Channels play super important role in Go Programming Language.
// Their main responsibility is to allow different go routines to
// communicate with each other by sharing memory through communication.
// 'Effective Go' specification strongly encourages us to use this approach
// rather then typical workflow with Mutexes.

func main() {
	// Use make() to build unbuffered channel. Unbuffered means
	// we cannot use scheduling, as well, as allow first go routine
	// to continue execution. First go routine will stops execution
	// at this point and waits until another go routine takes value
	// out of the channel.
	c := make(chan int)

	// Use channel <- to append data to the channel
	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
	}()

	// Use <- channel to receive data from the channel
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
	}()

	time.Sleep(time.Second)
}