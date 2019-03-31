package main

import (
	"fmt"
	"sort"
)

type person struct {
	Name string
	Age  int
}

type people []person

func (p people) Len() int           { return len(p) }
func (p people) Less(i, j int) bool { return p[i].Age < p[j].Age }
func (p people) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func main() {
	students := people{
		person{"Adam", 19},
		person{"Kelly", 21},
		person{"Al", 23},
	}

	fmt.Println("Students before sorting by age: ", students)

	sort.Sort(students)

	fmt.Println("Students after sorting by age: ", students)
}
