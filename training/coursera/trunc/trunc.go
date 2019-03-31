package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type input struct {
	stream string
}

func newInput(s string) *input {
	return &input{s}
}

func (i input) Exit() bool {
	return string(i.stream) == "exit"
}

func (i input) Float() (float64, error) {
	return strconv.ParseFloat(i.stream, 64)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	repeat := true

	for {
		fmt.Println("Type 'exit' to quit, or enter valid floating point number:")

		for scanner.Scan() {

			input := newInput(scanner.Text())

			if ok := input.Exit(); ok {
				fmt.Println("Program terminates...")
				repeat = false
				break
			}

			float, err := input.Float()

			if err != nil {
				fmt.Println("Invalid characters detected.")
				repeat = true
				break
			}

			fmt.Printf("Truncated from %s to %d\n", input.stream, int(float))
			break
		}

		if !repeat {
			break
		}
	}
}
