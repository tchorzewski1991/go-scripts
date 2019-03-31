package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := incrementor("C1: ", 6)
	c2 := incrementor("C2: ", 1)
	c3 := puller(c1)
	c4 := puller(c2)

	fmt.Println("Final counter: ", <-c3+<-c4)
}

func incrementor(str string, delay int) chan int {
	out := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			out <- 1
			fmt.Println(str, i)
			time.Sleep(time.Duration(delay) * time.Second)
		}
		close(out)
	}()
	return out
}

func puller(c chan int) chan int {
	out := make(chan int)
	go func() {
		var sum int
		for n := range c {
			sum += n
			fmt.Println("Sum: ", sum)
		}
		out <- sum
		close(out)
	}()
	return out
}
