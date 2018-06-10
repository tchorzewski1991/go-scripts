package main

import (
	"encoding/json"
	"fmt"
)

// We can manipulate which fields will be exported into
// JSON with lower / upper case first letter of the field
// name. Note as well how we can use feature called 'tags'
// in Golang. Basically tags allow for translating any
// field within our struct to some other alias name.
type Person struct {
	Name       string
	Surname    string
	occupation string
	Age        int    `json:"Years"`
}

func main() {

	// json.Marshal() returns a byte slice. It could be easily converted
	// into string, as string is nothing more than slice of bytes.

	person := Person{
		Name:       "Joe",
		Surname:    "Doe",
		occupation: "Programmer",
		Age:        21,
	}

	encoded, _ := json.Marshal(person)

	fmt.Println(encoded)
	fmt.Println(string(encoded))


	// json.Unmarshal() expects slice of bytes from raw JSON string
	// and an interface to fill it up with decoded values. Our interface
	// will be pointer to Person.

	var personPtr *Person

	rawJSONString  := `{"Name":"Joe", "Surname":"Doe", "Years":23}`

	rawBytes := []byte(rawJSONString)

	json.Unmarshal(rawBytes, &personPtr)

	fmt.Println(personPtr.Name)
	fmt.Println(personPtr.Surname)
	fmt.Println(personPtr.Age)
}
