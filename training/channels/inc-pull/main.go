package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("Starting main goroutine... \n")

	fmt.Println("Call incrementor() \n")
	c := incrementor()

	time.Sleep(3e9)

	fmt.Println("Call puller() \n")
	cSum := puller(c)

	fmt.Println("Waiting for cSum ...")
	for n := range cSum {
		fmt.Println("Pulling from cSum channel: \n")
		fmt.Println(n)
	}
}

func puller(c chan int) chan int {
	out := make(chan int)
	fmt.Println("Puller: before goroutine \n")

	go func() {
		var sum int
		for n := range c {
			fmt.Printf("Puller: ranging over c: %d \n", n)
			time.Sleep(1e9)
			sum += n
		}
		fmt.Println("Puller: before sending to out channel...")
		out <- sum
		fmt.Println("Puller: after sending to out channel...")

		fmt.Println("Puller: before closing out channel...")
		close(out)
		fmt.Println("Puller: after closing out channel...")
	}()
	fmt.Println("Puller: after goroutine \n")
	return out
}

func incrementor() chan int {
	out := make(chan int)
	fmt.Println("Incrementor: before goroutine \n")
	go func() {
		for i := 0; i < 5; i++ {
			out <- i
		}
		close(out)
	}()
	fmt.Println("Incrementor: after goroutine \n")
	return out
}
