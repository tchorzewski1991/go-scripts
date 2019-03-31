package main

import "fmt"

func main() {
	// Semaphore is simply a variable. It could be understood as an
	// abstract data type used to control access to a common resource
	// by multiple processes in concurrent systems. The main responsibility
	// of semaphores is to achieve process synchronization in the right way
	// (keep track over avoiding race conditions etc.)

	c := make(chan int)
	done := make(chan bool)

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		done <- true
	}()

	go func() {
		for i := 10; i < 20; i++ {
			c <- i
		}
		done <- true
	}()

	go func() {
		<-done
		<-done
		close(c)
	}()

	for n := range c {
		fmt.Println(n)
	}
}
