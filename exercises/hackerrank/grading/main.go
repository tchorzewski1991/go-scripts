// https://www.hackerrank.com/challenges/grading/problem

package main

import (
	"fmt"
)

const y = 5

func main() {
	grades := []int32{73, 67, 38, 33}

	grades = gradingStudents(grades)

	fmt.Println(grades)
}

func gradingStudents(grades []int32) []int32 {
	var r []int32

	for _, x := range grades {

		if x >= 38 && x%y >= 3 {
			r = append(r, x-(x%y)+y)
		} else {
			r = append(r, x)
		}
	}

	return r
}
