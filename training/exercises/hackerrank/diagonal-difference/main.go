package main

import (
	"fmt"
	"math/rand"
	"math"
)

func main() {
	var mtx [][]int32

	mtx = simpleMatrix(3, 3,  10, mtx)

	fmt.Println(mtx)
	fmt.Println(diagonalDifference(mtx))
}

// Fills up n x m dimensional matrix with random numbers from range 0 up to r.
// Notice, there is no random number seed generation, so it probably always
// fills up matrix with the same numbers.
func simpleMatrix(n int32, m int32, r int32, mtx [][]int32) [][]int32 {

	for i := 0; i < int(n); i++ {
		inner := make([]int32, m)

		for j := 0; j < int(m); j++ { inner[j] = rand.Int31n(r) }

		mtx = append(mtx, inner)
	}

	return mtx
}

// Computes absolute difference between each of diagonal sums.
func diagonalDifference(arr [][]int32) int32 {
	return func(arr [][]int32) int32 {
		var diff int32

		k := len(arr)

		for i := 0; i < k; i++ {
			diff += (arr[i][i] - arr[i][k - (1 + i)])
		}

		return int32(math.Abs(float64(diff)))
	}(arr)
}


