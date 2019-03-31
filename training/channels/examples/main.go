package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)

	go func() {
		fmt.Println("Receiving from chanel: \n", <-ch)
	}()

	fmt.Println("Sending to chanel -> 2")
	ch <- 2

	time.Sleep(2e9)
}
