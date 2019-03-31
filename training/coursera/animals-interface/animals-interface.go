package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type animal interface {
	eat()
	move()
	speak()
}

type cow struct {
	name       string
	food       string
	locomotion string
	noise      string
}

func (c cow) eat()   { fmt.Println(c.food) }
func (c cow) move()  { fmt.Println(c.locomotion) }
func (c cow) speak() { fmt.Println(c.noise) }

type bird struct {
	name       string
	food       string
	locomotion string
	noise      string
}

func (b bird) eat()   { fmt.Println(b.food) }
func (b bird) move()  { fmt.Println(b.locomotion) }
func (b bird) speak() { fmt.Println(b.noise) }

type snake struct {
	Name       string
	food       string
	locomotion string
	noise      string
}

func (s snake) eat()   { fmt.Println(s.food) }
func (s snake) move()  { fmt.Println(s.locomotion) }
func (s snake) speak() { fmt.Println(s.noise) }

func printProperty(animal animal, property string) {
	switch property {
	case "eat":
		animal.eat()
	case "move":
		animal.move()
	case "speak":
		animal.speak()
	default:
		{
			notFound("Property", property)
		}
	}
}

func newAnimal(name, _type string) (animal, bool) {
	switch _type {
	case "cow":
		return newCow(name), true
	case "bird":
		return newBird(name), true
	case "snake":
		return newSnake(name), true
	default:
		{
			return nil, false
		}
	}
}

func newCow(name string) *cow {
	return &cow{name, "grass", "walk", "moo"}
}

func newSnake(name string) *snake {
	return &snake{name, "mice", "slither", "hsss"}
}

func newBird(name string) *bird {
	return &bird{name, "worms", "fly", "peep"}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	repeat := true

	var animals = map[string]animal{}

	for {
		fmt.Println("Type: 'exit' to terminate program.")
		fmt.Println("Type: 'newanimal :animal_name :animal_type', to add new animal.")
		fmt.Println("Type: 'query :animal_name :animal_property', to get requested data.")

		for scanner.Scan() {
			input := scanner.Text()

			if input == "exit" {
				fmt.Println("Program terminates...")
				repeat = false
				break
			}

			parameters := strings.Split(input, " ")

			if len(parameters) != 3 {
				fmt.Println("Invalid input format. Please try again.")
				break
			}

			name := parameters[1]

			switch action := parameters[0]; action {
			case "newanimal":
				{
					typeOf := parameters[2]

					if _, found := animals[name]; found {
						alreadyFound("Animal name", name)
						break
					}

					animal, created := newAnimal(name, typeOf)

					if !created {
						notFound("Animal type", typeOf)
						break
					}

					animals[name] = animal

					fmt.Println("Created it!")
				}
			case "query":
				{
					prop := parameters[2]

					animal, found := animals[name]

					if !found {
						notFound("Animal name", name)
						break
					}

					printProperty(animal, prop)
				}
			default:
				{
					notFound("Action", action)
				}
			}

			break
		}

		if !repeat {
			break
		}
	}

	os.Exit(0)
}

func notFound(attrName, attrValue string) {
	fmt.Printf("%s %s not found. Please try again\n", attrName, attrValue)
}

func alreadyFound(attrName, attrValue string) {
	fmt.Printf("%s %s already exists. Please try another value\n", attrName, attrValue)
}
