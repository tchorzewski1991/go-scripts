//https://www.hackerrank.com/challenges/camelcase/problem
package main

import (
	"unicode"
	"fmt"
)

func main() {
	src := "oneTwoThree"

	var uppers int32

	for _, e := range src {
		if (unicode.IsUpper(e)) {
			uppers += 1
		}
	}

	fmt.Println(uppers)
}

