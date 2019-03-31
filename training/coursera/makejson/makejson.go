package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	name := scanAttribute("name")
	address := scanAttribute("address")
	composite := map[string]string{"name": name, "address": address}

	sb, err := json.Marshal(composite)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(sb))
}

func scanAttribute(attrName string) (attr string) {
	fmt.Printf("Enter %s: \n", attrName)
	_, err := fmt.Scan(&attr)
	if err != nil {
		panic(err)
	}
	return attr
}
