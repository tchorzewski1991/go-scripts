//https://www.hackerrank.com/challenges/mini-max-sum/problem

package main

import (
	"fmt"
	"sort"
)

// One of the possible solutions to this problem is to
// declare custom slice type and extend it later with sort.Interface
// It will make it sortable.
type ext_arr []int32

//Methods required by sort.Interface to be satisfied.
func (c ext_arr) Len()          int  { return len(c) }
func (c ext_arr) Swap(i, j int)      { c[i], c[j] = c[j], c[i] }
func (c ext_arr) Less(i, j int) bool { return c[i] < c[j] }

func main() {
	// Setup slice with custom elements order
	arr := []int32{1,4,3,5,7}
	miniMaxSum(arr)
}

// Finds the minimum and maximum values that can be calculated by summing
// exactly len(arr) - 1 of the len(arr) integers
func miniMaxSum(arr ext_arr) {
	var min, max int64
	var k = len(arr) - 1

	sort.Sort(arr)

	for i := 0; i < k; i++ {
		min += int64(arr[i]); max += int64(arr[k - i])
	}

	fmt.Println(min, max)
}
