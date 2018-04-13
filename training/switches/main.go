package main

import "fmt"

func main()  {

    // Switching between types is possible even if we must
    // to decide what the type is during the execution time.
    checkType := func (i interface{}) {
        switch t := i.(type) {
        default:
            fmt.Printf("Unexpected type: %T\n", t)
        case bool:
            fmt.Printf("Bool type: %T\n", t)
        case string:
            fmt.Printf("String type: %T\n", t)
        }
    }

    checkType("string")
    checkType(true)
	checkType(12)
}


