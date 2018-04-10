package main

import "fmt"

const breakPoint int = 10

func main() {

	var x int

	// Declaration of function expression that will be used as
	// conditional statement within for loop. Notice variable x
	// is visible on anonymous function block and do not need to be
	// provided explicitly. Default value for x will be resolved
	// automatically.
	condition := func () bool {
		return x < breakPoint
	}

	// Go do not have 'do - while' kind of loops. Instead we can
	// use more 'C' style loops like the one below.
	for condition() {
		fmt.Printf("x is : %d \n", x)
		x++
	}

	// Go allow for advanced control flow within for loops. Conditional
	// statement isn't required at all. Notice we get infinite loop
	// when control flow won't be explicitly provided in that case.
	x = 0

	for {
		x++

		if x % 2 == 0 { continue }

		fmt.Printf("x is : %d \n", x)

		if x >= 7 { break }
	}
}
