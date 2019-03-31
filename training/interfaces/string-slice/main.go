package main

import (
	"fmt"
	"sort"
)

type people []string

func main() {
	students := people{"Kelly", "John", "Al", "Cecil"}

	fmt.Println("Unsorted numbers: ", students)

	// sort.StringSlice is a type that defines implicitly
	// Interface interface. What does it mean is that all
	// methods required by sort.Sort() are already predefined.
	// We do need nothing to do on our side to provide sorting
	// functionality.
	sort.StringSlice(students).Sort()

	fmt.Println("Sorted students: ", students)
}
