package main

import "fmt"

func main() {

	// Printing number 42 in decimal and binary formats.

	// Printf() refers to variadic function which means it can take unlimited
	// number of arguments. When empty interface (...interface{}) is used as
	// one of the entries on arguments list it means arguments could be any type.
	fmt.Printf("%d - %b\n", 42, 42)

	printUTF(97, 122)
}

func printUTF(left int, right int) {
	for i := left; i <= right; i++ {
		fmt.Printf("%d - \t - %b - \t - %x - \t - %q \n", i, i, i, i)
	}
}
