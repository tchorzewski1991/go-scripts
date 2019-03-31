package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type input struct {
	stream string
}

func newInput(s string) *input {
	return &input{s}
}

func (i input) Exit() bool {
	return i.stream == "exit"
}

func (i input) hasAnySuffix(suffixes ...string) bool {
	hasAny := false
	for _, suffix := range suffixes {
		hasAny = strings.HasSuffix(i.stream, suffix)
		if hasAny {
			break
		}
	}
	return hasAny
}

func (i input) hasAnyPrefix(prefixes ...string) bool {
	hasAny := false
	for _, prefix := range prefixes {
		hasAny = strings.HasPrefix(i.stream, prefix)
		if hasAny {
			break
		}
	}
	return hasAny
}

func (i input) containsAny(chars ...string) bool {
	containsAny := false
	for _, char := range chars {
		containsAny = strings.Contains(i.stream, char)
		if containsAny {
			break
		}
	}
	return containsAny
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	repeat := true

	for {
		fmt.Println("Please type 'exit' to quit or enter new string:")

		for scanner.Scan() {
			input := newInput(scanner.Text())

			if ok := input.Exit(); ok {
				fmt.Println("Program terminates...")
				repeat = false
				break
			}

			hasAnyPrefix := input.hasAnyPrefix("i", "I")
			hasAnySuffix := input.hasAnySuffix("n", "N")
			containsAny := input.containsAny("a", "A")

			if hasAnySuffix && hasAnyPrefix && containsAny {
				fmt.Println("Found!")
			} else {
				fmt.Println("Not Found!")
			}

			break
		}

		if !repeat {
			break
		}
	}
}
