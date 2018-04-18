package main

import "fmt"

const newValue string = "new value"

func main() {

	// Prepend x with '&' causes fetching memory addresses of
	// that particular variable. We can notice that addresses of
	// local variables within both scopes (local and function) aren't
	// the same. In Golang we always refer to 'pass by value' approach,
	// rather than 'pass by reference' or 'pass by copy'. If we want
	// to change underlying value, we need to do pointer operations.
	var x string = "value"

	fmt.Println(&x)
	changeX(x)
	fmt.Println(x)

	// Using technique of referencing/de-referencing allows for
	// easy overwrites of underlying values. This is one of the
	// reasons why local to the outer scope variable y has been
	// changed from another scope.
	var y string = "value"

	fmt.Println(&y)
	changeY(&y)
	fmt.Println(y)
}

func changeX(x string) {
	fmt.Println(&x)
	x = newValue
}

func changeY(y *string) {
	fmt.Println(y)
	*y = newValue
}

