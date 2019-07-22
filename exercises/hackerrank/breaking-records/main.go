package main

// https://www.hackerrank.com/challenges/breaking-best-and-worst-records/problem

import "fmt"

func main() {
	scores := []int32{2, 3, 1, 4}
	fmt.Printf("%v\n", BreakingRecords(scores))
}

func BreakingRecords(scores []int32) []int32 {
	changes := []int32{0, 0}

	if len(scores) < 2 {
		return changes
	}

	min, max := scores[0], scores[0]

	for i := 0; i < len(scores); i++ {
		e := scores[i]

		if e < min {
			min = e
			changes[1]++
		}

		if e > max {
			max = e
			changes[0]++
		}
	}

	return changes
}
