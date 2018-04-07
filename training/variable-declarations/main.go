package main

import "fmt"

func main() {

	// Variable declaration

	var a string  // will return empty string by default
	var b int     // will return 0 by default
	var c float64 // will return 0 by default
	var d bool    // will return false by default

	fmt.Printf("%v \n", a)
	fmt.Printf("%v \n", b)
	fmt.Printf("%v \n", c)
	fmt.Printf("%v \n", d)

	// Variable assignment + initialization

	e := 10
	fmt.Printf("%v \n", e)
}
