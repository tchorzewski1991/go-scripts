package main

import "fmt"

// Structs in Golang are described as a sequences of fields
// where each field contains name and a type, or where field
// is just type-less promotion to another struct.

// Promotion means that each of the fields from inner struct
// will become injected (accessible) for outer parent struct.

// There is no classes and instantiating instances of those classes
// in Go. Instead, Golang offers ability to define custom types.
// Instantiation have been replaced with assigning values to those
// types.

type address struct {
	country  string
	city     string
	street   string
	postCode int
}

type person struct {
	firstName string
	lastName  string
	age       int
	address
}

func (p person) fullName() string {
	return p.firstName + " " + p.lastName
}

// Structs are reference types which means there is dependent
// underlying data structure under the hood. When we want to
// mutate state of the given struct we can use pointer manipulations
// to do that. Very common solution is to do this with receiver
// function.

func (p *person) changeAge(newAge int) {
	p.age = newAge
}

func main() {

	person := person{
		firstName: "Joe",
		lastName:  "Doe",
		age:       20,
		address: address{
			country:  "Country",
			city:     "City",
			street:   "Street",
			postCode: 666,
		},
	}

	fmt.Printf("Person name: %s \n", person.firstName)
	fmt.Printf("Person city: %s \n", person.address.city)
	fmt.Printf("Person age:  %d \n", person.age)

	person.changeAge(21)

	// Receiver functions are called also idiomatically 'methods'.
	fmt.Printf("Person full name: %s \n", person.fullName())
	fmt.Printf("Person new age:   %d \n", person.age)

}
