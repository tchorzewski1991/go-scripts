package main

import "fmt"

// Variable x is declared and assigned on the 'package' level
// scope. This scope encloses both functions and allows for
// accessing x easily. Notice use of '++' operator which
// increments value of variable x without any problems.
// It means that we're accessing direct reference to that value,
// rather than a shallow copy, as its always happens in case of
// functions. See examples below.

var x int = 0

func increment() int  {
	x++
	return x
}

func main() {
	fmt.Println("x is: ", x)
	increment()

	fmt.Println("x is: ", x)
	increment()

	fmt.Println("x is: ", x)
	increment()
}
