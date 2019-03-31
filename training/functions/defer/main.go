package main

import "fmt"

// Golang allows you to specify your own type easily.
// It might be understood as something like monkey patching
// existing predefined types, but without truly abusing
// their original code base and state.
type mapType map[string]string

// Golang allows for specifying receiver functions. This functions
// refers to bounding (defining) behavior with specific type.
// In other words it could be understood as defining function on
// self, where self is the given type.
func (mt mapType) printValues() {
	for key, value := range mt {
		fmt.Printf("key: %s, value %s \n", key, value)
	}
}

func (mt mapType) callback() {
	fmt.Println("callback called!")
}

func main() {
	myMap := mapType{
		"name":    "Joe",
		"surname": "Doe",
	}

	myMap.printValues()

	// Deferring in Go is technique useful especially when we need to
	// be sure that some behavior will be invoked as a callback. In
	// declaration below we ensure that myMap.callback() will be executed
	// just before our main() function ends.
	defer myMap.callback()
}
