package main

import (
	"bufio"
	"fmt"
	"github.com/pkg/errors"
	"os"
	"sort"
	"strconv"
	"strings"
)

type aggregate struct {
	arrays [][]int
}

func newAggregate(ints []int) *aggregate {
	temp := make([][]int, 4)

	for i := 0; i < len(ints); i++ {
		temp[i%4] = append(temp[i%4], ints[i])
	}

	return &aggregate{temp}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var sequence []int
	var err error

	fmt.Println("Enter sequence of integers separated by space: ")

	for scanner.Scan() {
		sequence, err = parseInts(scanner.Text())
		break
	}

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	aggregate := newAggregate(sequence)

	var result []int

	for _, arr := range aggregate.arrays {
		result = append(result, <-sortArray(arr)...)
	}

	sort.Ints(result)

	fmt.Println("Final result:     ", result)
	os.Exit(0)
}

func parseInts(line string) ([]int, error) {
	var result []int

	for _, s := range strings.Split(line, " ") {

		if len(s) == 0 {
			continue
		}

		i, err := strconv.Atoi(strings.TrimSpace(s))

		if err != nil {
			return nil, errors.New(fmt.Sprint(s, " is not a valid character"))
		}

		result = append(result, i)
	}

	return result, nil
}

func sortArray(arr []int) chan []int {
	out := make(chan []int)

	go func() {
		sort.Ints(arr)
		fmt.Println("Sorted sub-array: ", arr)
		out <- arr
		close(out)
	}()

	return out
}
