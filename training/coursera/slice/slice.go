package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	result := make([]int, 0)
	repeat := true

	for {
		fmt.Println("Please type 'X' to quit or enter new integer:")

		for scanner.Scan() {
			input := scanner.Text()

			if input == "X" {
				fmt.Println("Program terminates.")
				repeat = false
				break
			}

			i, err := strconv.Atoi(input)

			if err != nil {
				fmt.Println("Invalid characters detected. Please provide valid integer.")
				break
			}

			result = append(result, i)
			sort.Ints(result)

			fmt.Println(result)

			break
		}

		if !repeat {
			break
		}
	}
}
