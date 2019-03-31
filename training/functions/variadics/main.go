package main

import "fmt"

func main() {

	// Variadic function can be called with unlimited number of arguments.
	// Notice nums will be referenced inside function body as a slice of
	// floats64's. ([]float64). If types of arguments that we have are
	// unknown we ca use ...interface{} instead.
	average := func(nums ...float64) float64 {
		var total float64

		for _, n := range nums {
			total += n
		}

		return total / float64(len(nums))
	}

	avg := average(12, 13, 14)

	// Should print 13. It refers to the average of numbers specified
	// as arguments list on the average() function call above.
	fmt.Println("Average is: ", avg)

	// Variadic arguments refers to approach where our function expects
	// unlimited number of arguments (could be with resolved type) but
	// instead of providing arguments one-by-one we want to 'unpack'
	// existing data structure - for example slice.
	accumulate := func(nums ...int) int {
		var accumulator int

		for _, n := range nums {
			accumulator += n
		}

		return accumulator
	}

	numbers := []int{1, 2, 3}

	sum := accumulate(numbers...)

	// Should print 6. It refers to the sum of all values from slice of int32's.
	// Notice how accumulate() argument 'numbers' are appended with '...'.
	// It basically means - 'take those slice of int's and append one by one
	// as a new argument to variadic function call.
	fmt.Println("Sum is: ", sum)
}
