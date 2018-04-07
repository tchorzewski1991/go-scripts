package main

// Explicit package import allows for using globally visible
// variables / functions by direct reference through package name.
import (
	"github.com/tchorzewski1991/go-scripts/training/scopes/support"
	"fmt"
)

// Declaration of custom string type.
type supportConf string

func main() {

	// Value of support.Conf can be easily re-referenced with custom type.
	str := supportConf(support.Conf)
	fmt.Println(str)

	// Value of support.Conf is visible in main() block scope.
	fmt.Println(support.Conf)

	// Value of support.Conf is visible in outer scope.
	printSupportConf()
}

func printSupportConf() {
	fmt.Println(support.Conf)
}
