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

	// It's impossible to enforce the same behavior with local variable
	// given to the function as an argument. Go programming language is
	// 'pass by value' language. If we provide 'VALUE' type argument to
	// the function call, copy of that variable will stored on new empty
	// memory slot.

	x = 0

	// Variable won't be successfully incremented as we operate on copy.
	incrementByOne(x)
	fmt.Println("x is: ", x)

	// Variable will be successfully incremented because of pointers manipulations.
	incrementPointerByOne(&x)
	fmt.Println("x is: ", x)
}

func incrementByOne(x int)  {
	x++
}

func incrementPointerByOne(x *int)  {
	*x++
}
