package main

import (
	"fmt"
	"time"
)

type square struct {
	a float64
}

func (s square) area() float64 {
	return s.a * s.a
}

type rectangle struct {
	a float64
	b float64
}

func (r rectangle) area() float64 {
	return r.a * r.b
}

// Interfaces in Go Programming Language could be compared to a form
// of polymorphism. It is an OOP pattern where basically two
// different entities provide the functionality, but about underlying
// implementation decides given data type.

type shape interface {
	area() float64
}

func displayArea(s shape) {
	fmt.Println("Area of shape is: ", s.area())
}


func main() {
	r := rectangle{1.0, 1.5}
	s := square{1.5}

	displayArea(r)
	displayArea(s)

	time.Sleep(time.Second)
}
