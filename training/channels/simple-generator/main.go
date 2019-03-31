package main

import (
	"fmt"
	"os"
	"time"
)

func f1(in chan int) {
	fmt.Println("Listening for value from channel: \n")

	if v, ok := <-in; ok {
		fmt.Printf("Value received: %d \n", v)
	}

	fmt.Println("Function will terminate after 1 sec.")
	time.Sleep(time.Duration(1 * time.Second))
}

func main() {
	//ch := make(chan int)
	//done := make(chan bool)

	fmt.Println("Starting main goroutine ...")

	out := make(chan int)
	go f1(out)

	time.Sleep(time.Duration(3 * time.Second))

	fmt.Println("Sending value through channel")

	out <- 2

	//
	//go pump(ch, done)

	//go func() {
	//	<-done
	//	close(ch)
	//}()
	//
	//for i := range ch {
	//	fmt.Printf("i: %d \n", i)
	//}

	//for i := range ch {
	//	fmt.Printf("i: %d \n", i)
	//}

	//for {
	//	fmt.Printf("i: %d \n", <- ch)
	//}

	time.Sleep(time.Duration(1 * time.Second))
	os.Exit(1)
}

//func pump(ch chan int, done chan bool) {
//	for i := 1; i <= 10; i++ {
//		ch <- i
//	}
//	done <- true
//}
