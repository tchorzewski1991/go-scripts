package main

import (
	"fmt"
)

func main() {
	c := make(chan int)

	// Close() prevents from adding more values onto the channel.
	// We can still receive values from the channel, but we
	// can't put anymore values in it.
	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		close(c)
	}()

	// Ranges automatically fetches values out of the channel one
	// after another. Notice how there is no more need for new
	// go routine to be up and running. The flow of execution is
	// as follows:
	// 1. We initialize new, empty, unbuffered channel
	// 2. We are scheduling fresh go routine to be up and
	//    running on it own.
	// 3. Code halts at for - range clause and waits until some value
	//    will be received out of the channel. There is no more need for
	//    using time.Sleep() at the end to prevent main from self terminating.
	for n := range c {
		fmt.Println(n)
	}
}
