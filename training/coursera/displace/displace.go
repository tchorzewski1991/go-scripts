package main

import (
	"fmt"
	"github.com/pkg/errors"
	"math"
	"os"
	"strconv"
)

func main() {
	var a, v0, s0 float64

	initialAttributes := map[string]float64{
		"acceleration": a,
		"velocity":     v0,
		"displacement": s0,
	}

	for name := range initialAttributes {
		initialAttributes[name] = prepareAttribute(name)
	}

	displaceFn := genDisplaceFn(
		initialAttributes["acceleration"],
		initialAttributes["velocity"],
		initialAttributes["displacement"],
	)

	var time float64
	time = prepareAttribute("time")
	fmt.Printf("Displacement equals %.3f \n", displaceFn(time))

}

func genDisplaceFn(a float64, v0 float64, s0 float64) func(t float64) float64 {
	return func(t float64) float64 {
		return 0.5*a*math.Pow(t, 2.0) + v0*t + s0
	}
}

func prepareAttribute(attrName string) float64 {
	input, err := scanInput(attrName)
	if err != nil {
		handleError(err)
	}

	number, err := convertToFloat64(&input)
	if err != nil {
		handleError(err)
	}

	if number < 0 {
		handleError(errors.New("attribute cannot be less than 0"))
	}

	return number
}

func scanInput(attrName string) (input string, err error) {
	fmt.Printf("Please enter value for %s: \n", attrName)
	_, err = fmt.Scan(&input)
	if err != nil {
		return "", errors.New(err.Error())
	}
	return input, nil
}

func convertToFloat64(inputPtr *string) (number float64, err error) {
	number, err = strconv.ParseFloat(*inputPtr, 64)
	if err != nil {
		return 0, errors.New(err.Error())
	}
	return number, nil
}

func handleError(err error) {
	fmt.Println("An error has occured. ", err)
	os.Exit(1)
}
