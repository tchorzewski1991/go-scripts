package main // https://www.hackerrank.com/challenges/kangaroo/problem
import "fmt"

func main() {
	var x1, v1, x2, v2 int32

	x1, v1 = 4523, 8092
	x2, v2 = 9419, 8076

	fmt.Println(Kangaroo(x1, v1, x2, v2))
}

func Kangaroo(x1 int32, v1 int32, x2 int32, v2 int32) string {
	var verdict string

	sizeLimiter := func(values ...int32) bool {
		var result bool
		for _, v := range values {
			if v > 10000 {
				result = true
				break
			}
		}
		return result
	}

	for {
		if ok := sizeLimiter(x1, v1, x2, v2); ok {
			verdict = "NO"
			break
		}

		if x1 < x2 && v1 < v2 {
			verdict = "NO"
			break
		}

		if x1 == x2 {
			verdict = "YES"
			break
		}

		x1 = x1 + v1
		x2 = x2 + v2
	}

	return verdict
}
