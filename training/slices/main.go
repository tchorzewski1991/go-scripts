package main

import "fmt"

func main() {
	// Slice is a data structure build on top of an array. It is
	// a list of elements of the same type. It contains length
	// and capacity. It's pure reference type and it needs to be
	// considered wisely, especially when it comes to initialization.

	// The most appropriate way to create a slice is to use 'make'
	// build-in function. We provide length and capacity initially so
	// there is no error while accessing 0 index position. Error will
	// be raised when we try to access index position greater then actual
	// slice length.

	fromMake := make([]int, 1, 1)
	fromMake[0] = 42
	fmt.Println("There is no error while accessing 0 index position", fromMake[0])

	// This line will make an out of range index error.
	//fromMake[1] = 43

	// We can use built-in 'append' function to expand slice capacity
	// and length. This is the appropriate way to do this.
	fromMake = append(fromMake, 43)
	fmt.Println("There is no error while accessing 1 index position", fromMake[1])

	// When we don't use built-in make function we need to take into
	// account that our declared slice will be pointing to nil. This
	// is default value for slice data structure in Golang.
	var slice []int
	fmt.Println("Empty slice with default value of nil: ", slice)

	//fmt.Println("We can't access 0 index position as it isn't exist", slice[0])
}
