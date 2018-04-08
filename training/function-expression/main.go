package main

import "fmt"

var x int

func main() {
	// Example of function expression. This is one of the way to define
	// function within another function block scope.
	// Notice: func () {} is anonymous function.
	increment := func () int {
		x++
		return x
	}

	fmt.Printf("%d \n", increment())
	fmt.Printf("%d \n", increment())
}
