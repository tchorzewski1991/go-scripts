package main

import "fmt"

func multiReturn(x int, y int) (i int, j int) {
	return y, x
}

func main() {

	// Leaving unused variables in Go is not allowed. It is
	// categorized as a code pollution and compiler won't let
	// you do that. To indicate we don't need some value
	// anymore, Go encourages to use special syntax or
	// technique called 'blank identifiers'. By default sign
	// that expresses the intent described above is a low dash.
	var x int
	x, _ = multiReturn(42, 43)

	fmt.Printf("x is : %d \n", x)
}
