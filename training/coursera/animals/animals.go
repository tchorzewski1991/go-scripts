package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type animal struct {
	food       string
	locomotion string
	noise      string
}

func (a animal) Eat()   { fmt.Println(a.food) }
func (a animal) Move()  { fmt.Println(a.locomotion) }
func (a animal) Speak() { fmt.Println(a.noise) }

func main() {
	animals := map[string]animal{
		"cow":   {"grass", "walk", "moo"},
		"bird":  {"worms", "fly", "peep"},
		"snake": {"mice", "slither", "hss"},
	}

	repeat := true
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("Predefined animals:    cow/bird/snake")
		fmt.Println("Predefined properties: eat/move/speak")
		fmt.Println("Type: :animal :property to fetch information")
		fmt.Println("Type: 'exit' to terminate program")

		for scanner.Scan() {
			input := scanner.Text()

			if input == "exit" {
				repeat = false
				break
			}

			tokens := strings.Split(input, " ")

			if len(tokens) != 2 {
				fmt.Println("Invalid input format. Please use :animal :property instead\n")
				break
			}

			if animal, ok := animals[tokens[0]]; ok {
				switch tokens[1] {
				case "eat":
					animal.Eat()
					fmt.Println()
				case "move":
					animal.Move()
					fmt.Println()
				case "speak":
					animal.Speak()
					fmt.Println()
				default:
					{
						fmt.Println("Property not found. Choose one of: eat/move/speak\n")
						break
					}
				}
			} else {
				fmt.Println("Animal not found. Choose one of: cow/bird/snake\n")
				break
			}

			break
		}

		if !repeat {
			break
		}
	}

	fmt.Println("Program terminates...")
	os.Exit(0)
}
