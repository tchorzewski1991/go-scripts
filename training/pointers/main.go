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
}
