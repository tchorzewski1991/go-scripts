package main

import "fmt"

func main() {

	// Fetch memory address by prepending variable with &
	x := 42
	fmt.Printf("%p \n", &x)
}
