package main

import (
	"fmt"
	"os"
	"runtime"
	"sync"
	"time"
)

var wg sync.WaitGroup
var mt sync.Mutex

func main() {
	fmt.Println("Program starts...\n")

	runtime.GOMAXPROCS(2)

	log := func(msg string) {
		fmt.Println(msg, "letters goroutine has finished\n")
	}

	wg.Add(2)
	{
		go func() {
			defer func() { log("Lowercase"); wg.Done() }()
			for count := 0; count < 3; count++ {
				for char := 'a'; char < 'a'+26; char++ {
					fmt.Printf("%c ", char)
				}
			}
		}()

		go func() {
			defer func() { log("Uppercase"); wg.Done() }()
			for count := 0; count < 3; count++ {
				for char := 'A'; char < 'A'+26; char++ {
					fmt.Printf("%c ", char)
				}
			}
		}()
	}

	fmt.Println("\nWaiting to finish...\n")
	wg.Wait()

	fmt.Println("Program terminates in 2 sec...\n")
	time.Sleep(time.Duration(2 * time.Second))
	fmt.Println("Program terminated.")
	os.Exit(0)
}
