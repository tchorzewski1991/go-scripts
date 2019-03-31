package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type name struct {
	fname string
	lname string
}

func main() {
	var composite []name
	var fileName string

	fmt.Println("Enter file name: ")
	_, err := fmt.Scan(&fileName)

	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("No such file or directory.")
		os.Exit(1)
	}

	sc := bufio.NewScanner(file)

	for sc.Scan() {
		line := sc.Text()

		if len(line) == 0 {
			continue
		}

		initials := strings.Split(line, " ")

		composite = append(composite, name{initials[0], initials[1]})
	}

	for _, c := range composite {
		fmt.Printf("%s, %s \n", c.fname, c.lname)
	}
}
