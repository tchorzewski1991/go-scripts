//https://www.hackerrank.com/challenges/a-very-big-sum

package main

import "fmt"

func main() {
	arr := []int64{1000000001, 1000000002, 1000000003, 1000000004, 1000000005}
	fmt.Println(aVeryBigSum(arr))
}

func aVeryBigSum(arr []int64) int64 {
	var sum int64
	for _, e := range arr { sum += e }
	return sum
}