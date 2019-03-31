package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int32{2, 3, 1, 3}
	tcc := birthdayCakeCandles(arr)
	fmt.Println("Number of tallest cake candles: ", tcc)
}

func birthdayCakeCandles(arr []int32) int32 {
	var temp []int

	var tcc int32 = 1

	for _, e := range arr {
		temp = append(temp, int(e))
	}

	sort.Ints(temp)

	k := len(temp) - 1

	for i := 0; i < k; i++ {
		if temp[k-i] != temp[k-i-1] {
			break
		}
		tcc++
	}

	return tcc
}
