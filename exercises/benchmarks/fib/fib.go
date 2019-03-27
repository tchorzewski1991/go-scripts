package fib

// RecursiveFib is recursive implementation of the fibonacci algorithm
func RecursiveFib(n int) int {
	if n < 2 {
		return n
	}
	return RecursiveFib(n-2) + RecursiveFib(n-1)
}

// IterativeFib is iterative implementation of the fibonacci algorithm
func IterativeFib(n int) int {
	res := make([]int, n+1)
	res[0], res[1] = 0, 1

	for i := 2; i < n+1; i++ {
		res[i] = res[i-2] + res[i-1]
	}
	return res[n]
}
