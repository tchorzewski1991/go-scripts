package main

import "fmt"

func main() {

	valueToCheck := 3

	// Golang allows for attaching name to the value returned by function.
	// This is not quite frequently used pattern, but could be interesting
	// in some cases.
	divideableByTwo := func(i int) (result bool) {
		result = i%2 == 0
		return
	}

	// Should return false, as 3 is not divideable by 2.
	fmt.Println(divideableByTwo(valueToCheck))
}
