package main

import "fmt"

func main() {

	// Map is a key - value data structure. It's unordered and pure
	// 'dictionary' like reference type. Default value for map is
	// nil, so it needs to be considered wisely during initialization.
	// Each key and value are statically typed and cannot be
	// represented by other types.
	wordsCounter := func (words ...string) map[string]int {
		result := make(map[string]int)

		for _, word := range words {
			result[word]++
		}

		return result
	}

	result := wordsCounter("a", "b", "a", "c")
	fmt.Printf("Words count: %+v \n", result)

	// You won't be able to make assignment when map is declared in
	// the way presented below, as it points to nil value

	//var nilMap map[string]string
	//nilMap["key"] = "value"
	//fmt.Printf("Assignment won't happen. An error will occure: %v \n", makeMap)

	// Map declaration with build-in make() function won't produce map
	// that points to nil reference. It allows you to assign any key - value
	// pair with correct types.
	var makeMap = make(map[string]string)

	makeMap["firstKey"] = "firstValue"
	makeMap["lastKey"]  = "lastKey"

	fmt.Printf("Assignment will happen. No errors encountered: %v \n", makeMap)

	// Composite literal is the last technique of declaring new maps that won't
	// points under the hood to nil reference.
	compositeLiteralMap := map[string]string{
		"firstKey": "firstValue",
		"lastKey":  "lastValue",
	}

	fmt.Printf("Assignment will happen. No errors encountered: %v \n", compositeLiteralMap)
}

