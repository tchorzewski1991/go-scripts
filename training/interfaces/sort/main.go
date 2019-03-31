package main

import (
	"fmt"
	"sort"
)

// To implement an interface means to provide a specific
// set of behavior that is required by that interface.
// To use properly sort.Sort() function we need to provide
// as an argument something that implements Interface interface.
// Specifically to allow that we need to setup three additional
// methods on our people concrete type. This methods are Len(),
// Less(), Swap(). We should precisely adjust their behavior
// to our needs. In this simple example we will compare every
// string within the slice by it first letter.

type people []string

func (p people) Len() int {
	return len(p)
}

func (p people) Less(i, j int) bool {
	return p[i] < p[j]
}

func (p people) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func main() {
	students := people{"Mark", "Al", "Anthony", "Karl", "Mery"}

	fmt.Println("Unsorted students: ", students)

	sort.Sort(students)

	fmt.Println("Sorted students: ", students)
}
