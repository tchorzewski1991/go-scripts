package main

import (
	"fmt"
	"github.com/tchorzewski1991/go-scripts/training/pointers/converter"
)

func main() {

	// Access to memory address of the value referenced by variable could
	// be achieved easily by using & operator in front of the variable.
	x := 42
	fmt.Printf("%p \n", &x)

	// Example of using pointers. Notice that fmt.Scan() function
	// takes a pointer as an argument rather than value itself.
	// We want to feed exactly that particular memory slot with user
	// input.
	fmt.Print("Enter meters: ")

	var meters float64

	fmt.Scan(&meters)
	fmt.Printf("%v \n", converter.MetersToYards(meters))

	// Example of reference / dereference. Notice y variable is a type
	// of pointer to int. We directly assign to it memory address of
	// variable x, so we can easily overwrite this memory slot in the
	// future. Term '*y' means 'take value of y (memory address) and
	// assign to it something new.
	x = 43

	fmt.Printf("x is: %d \n", x)
	fmt.Printf("x address is %p \n", &x)

	var y = &x
	*y = 42

	fmt.Printf("x is: %d \n", x) // should be 42.
}
