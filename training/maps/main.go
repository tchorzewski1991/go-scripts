package main

import "fmt"

func main() {

	// Map is a key - value data structure. It's unordered and pure
	// 'dictionary' like reference type. Default value for map is
	// nil, so it needs to be considered wisely during initialization.
	// Each key and value are statically typed and cannot be
	// represented by other types.
	wordsCounter := func(words ...string) map[string]int {
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
	makeMap["lastKey"] = "lastKey"

	fmt.Printf("Assignment will happen. No errors encountered: %v \n", makeMap)

	// Composite literal is the last technique of declaring new maps that won't
	// points under the hood to nil reference.
	compositeLiteralMap := map[string]string{
		"firstKey": "firstValue",
		"lastKey":  "lastValue",
	}

	fmt.Printf("Assignment will happen. No errors encountered: %v \n", compositeLiteralMap)

	// We can easily remove any key - value pair from map with built-in delete() function.
	// Notice, how this approach uses standard look-up for existing key with 'for' loop.

	mapBeforeRemoveStandard := map[int]string{
		1: "one",
		2: "two",
	}

	fmt.Printf("Before key remove: %v \n", mapBeforeRemoveStandard)

	standardKeyRemover := func(k int, mapping map[int]string) map[int]string {
		for key := range mapping {
			if k == key {
				delete(mapping, key)
			}
		}

		return mapping
	}

	mapAfterRemoveStandard := standardKeyRemover(1, mapBeforeRemoveStandard)

	fmt.Printf("After key remove: %v \n", mapAfterRemoveStandard)

	// We can remove specific key - value pairs from map with more idiomatic approach
	// as well. We won't need any loop at all to achieve the same behavior.
	mapBeforeRemoveIdiomatic := map[int]string{
		1: "one",
		2: "two",
	}

	fmt.Printf("Before key remove: %v \n", mapBeforeRemoveIdiomatic)

	idiomaticKeyRemover := func(k int, mapping map[int]string) map[int]string {
		if _, exists := mapping[k]; exists {
			delete(mapping, k)
		}

		return mapping
	}

	mapAfterRemoveIdiomatic := idiomaticKeyRemover(1, mapBeforeRemoveIdiomatic)

	fmt.Printf("After key remove: %v \n", mapAfterRemoveIdiomatic)
}
