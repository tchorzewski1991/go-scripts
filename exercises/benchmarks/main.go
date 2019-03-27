package main

import (
	"fmt"
	"github.com/tchorzewski1991/go-scripts/exercises/benchmarks/fib"
)

func main() {
	fmt.Println("Fibonacci recursive call for 10: ", fib.RecursiveFib(10))
	fmt.Println("Fibonacci iterative call for 10: ", fib.IterativeFib(10))
}
