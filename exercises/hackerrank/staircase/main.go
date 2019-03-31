package main

import (
	"fmt"
	"strings"
)

func main() {
	n := int32(6)
	staircase(n)
}

// Prints staircase level using # sign where n refers to total number
// of levels to generate
func staircase(n int32) {
	for i := n - 1; i >= 0; i-- {
		lvl := newLevel(n)
		lvl = strings.Replace(lvl, "#", " ", int(i))
		fmt.Println(lvl)
	}
}

// Returns string of size n (new staircase level) filled with "#"
func newLevel(n int32) string {
	ss := make([]string, n)
	var i int32
	for i = 0; i < n; i++ {
		ss = append(ss, "#")
	}
	return strings.Join(ss, "")
}
