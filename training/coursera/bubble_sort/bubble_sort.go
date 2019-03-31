package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Enter sequence of integers separated by space: ")

	scanner := bufio.NewScanner(os.Stdin)

	var tokens []string
	var ints []int

	for scanner.Scan() {
		parseTokens(scanner.Text(), " ", &tokens)
		parseInts(tokens, &ints)
		bubbleSort(ints)
		break
	}

	fmt.Println(ints)
}

func parseTokens(line, sep string, tokens *[]string) {
	for _, s := range strings.Split(line, sep) {
		*tokens = append(*tokens, s)
	}
}

func parseInts(tokens []string, ints *[]int) {
	for _, token := range tokens {
		i, err := strconv.ParseInt(token, 0, 0)
		if err != nil {
			continue
		}
		*ints = append(*ints, int(i))
	}
}

func bubbleSort(ints []int) {
	l := len(ints)

	for i := 0; i < l; i++ {
		k := l - (i + 1)

		for j := 0; j < k; j++ {
			if ints[j] > ints[j+1] {
				swap(ints, j)
			}
		}
	}
}

func swap(temp []int, i int) {
	temp[i], temp[i+1] = temp[i+1], temp[i]
}
