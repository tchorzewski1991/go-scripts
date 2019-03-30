package main

import (
	"fmt"
	"strconv"
	"sync"
)

// Race conditions occur when two or more Goroutines try to
// use a piece of shared data at the same time. It is possible
// that code presented below will be executed multiple times.
// One way we could prevent a data race is to ensure that if
// one Goroutine is using the Balance variable, then all other
// Goroutines are prevented (or mutually excluded) from using
// it at the same time. This is where Mutex structure starts to
// shine.
// While one Goroutine holds the lock, all other Goroutines are
// prevented from executing any lines of code protected by the
// same mutex, and are forced to wait until the lock is yielded
// before they can proceed.

type currency struct {
	amount float64
	code   string
	mu     *sync.Mutex
}

func (c *currency) Add(i float64) {
	c.mu.Lock()
	c.amount += i
	c.mu.Unlock()
}

func (c *currency) Display() string {
	c.mu.Lock()
	defer c.mu.Unlock()
	return strconv.FormatFloat(c.amount, 'f', 2, 64)
}

var balance = currency{amount: 50.0, code: "GBP"}

func main() {
	fmt.Printf("Balance before addition: %s \n", balance.Display())

	balance.Add(12.0)

	fmt.Printf("Balance after addition: %s \n", balance.Display())
}
