package main

import "fmt"

func main() {

	sum := 0

	// Golang allows for variable assignment during  'if' conditional statement
	// declaration. Notice it is allowed to assign even anonymous function.
	// The scope of assigned variable is the 'if' block which prevents outer
	// scope from unnecessary code pollutions.

	for n := 0; n < 1000; n++ {
		if mod3 := func(i int) bool { return i%3 == 0 }(n); mod3 {
			sum += n
		}
	}

	// Should print sum of all numbers lower than 1000 and dividable by 3.
	fmt.Println("Sum is:", sum)

}
