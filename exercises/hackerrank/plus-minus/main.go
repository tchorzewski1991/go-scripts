package main

import "fmt"

// See https://www.hackerrank.com/challenges/plus-minus/problem for description.

func plusMinus(arr []int32) {
	var frPositive, frNegative, frZeros int32

	l := int32(len(arr))

	for _, e := range arr {
		switch {
		case e > 0:
			frPositive++
		case e < 0:
			frNegative++
		default:
			frZeros++
		}
	}

	fmt.Printf("%0.6f\n", getFraction(frPositive, l))
	fmt.Printf("%0.6f\n", getFraction(frNegative, l))
	fmt.Printf("%0.6f\n", getFraction(frZeros, l))
}

func getFraction(fr int32, l int32) float64 {
	return float64(fr) / float64(l)
}

func main() {
	arr := []int32{-1, 1, 0, 2, 1}
	plusMinus(arr)
}
