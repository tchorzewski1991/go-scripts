package main

import "fmt"

// See https://www.hackerrank.com/challenges/plus-minus/problem for description.

func plusMinus(arr []int32) {
	var frPositive, frNegative, frZeros int32

	arr_l := int32(len(arr))

	for _, e := range arr {
		switch {
		case e > 0:
			frPositive += 1
		case e < 0:
			frNegative += 1
		default:
			frZeros += 1
		}
	}

	fmt.Printf("%0.6f\n", getFraction(frPositive, arr_l))
	fmt.Printf("%0.6f\n", getFraction(frNegative, arr_l))
	fmt.Printf("%0.6f\n", getFraction(frZeros, arr_l))
}

func getFraction(fr int32, arr_l int32) float64 {
	return float64(fr) / float64(arr_l)
}

func main() {
	arr := []int32{-1, 1, 0, 2, 1}
	plusMinus(arr)
}
