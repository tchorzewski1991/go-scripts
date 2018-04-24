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
	fmt.Printf("Words count: %+v", result)
}
