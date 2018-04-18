package main

import "fmt"

// Golang offers functions as a type. We can pass function as an
// argument just like we do with strings or other types.
func filter(numbers []int, callback func(i int) bool) []int {
	result := []int{}

	for _, n := range numbers {
		if callback(n) {
			result = append(result, n)
		}
	}

	return result
}

func main() {
	result := filter([]int{1,2,3,4,5,6}, func(n int) bool {
		return n > 3
	})

	// Should return [4,5,6] as those are only numbers that satisfy
	// condition provided with anonymous function.
	fmt.Println("Result is: ", result)
}

