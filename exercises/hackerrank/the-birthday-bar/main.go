package main

// https://www.hackerrank.com/challenges/the-birthday-bar/problem

import (
	"fmt"
)

func main() {
	nums := []int32{}
	fmt.Printf("%v\n", Birthday(nums, 4, 2))
}

func Birthday(s []int32, d int32, m int32) int32 {
	var sums int32

	betweenFn := func(min, max, value int32) bool {
		return value < min || max < value
	}

	if betweenFn(1, 100, int32(len(s))) {
		return sums
	}
	if betweenFn(1, 31, d) {
		return sums
	}
	if betweenFn(1, 12, m) {
		return sums
	}
	if betweenFn(1, int32(len(s)), m) {
		return sums
	}

	i := int32(0)
	l := int32(len(s)) - m

	for ; i <= l; i++ {
		var temp int32

		for j := i; j < i+m; j++ {
			temp = temp + s[j]
		}

		if d == temp {
			sums++
		}
	}

	return sums
}
