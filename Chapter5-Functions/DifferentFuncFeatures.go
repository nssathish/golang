package main

import (
	"errors"
	"fmt"
)

//blank return types - never use these
func divAndReminder(numerator, denominator int) (result, reminder int, err error) {
	if denominator == 0 {
		err = errors.New("cannot divide by 0")
		return
	}
	result = numerator / denominator
	reminder = numerator % denominator
	err = nil

	return result, reminder, err
}

func multipleReturnValues() (int, int, error) {
	return 3, 0, errors.New("cannot divide by 0")
}

func someOne(guy person) {
	fmt.Println(guy.name, "is a person")
}

func division(numerator float64, denominator float64) float64 {
	if denominator == 0 {
		return 0
	}
	fmt.Println(variadicInputParameters(2, 1, 2, 3, 4, 5))
	return numerator / denominator
}

func variadicInputParameters(base int, vals ...int) []int {
	values := make([]int, 0, len(vals))
	for _, v := range vals {
		values = append(values, base+v)
	}

	fmt.Println("Parameterised: ", variadicInputParameters1(base, vals...)) // unpacking the slice ***
	return values
}

func variadicInputParameters1(base int, vals ...int) int {
	values := make([]int, 0, len(vals))
	for _, v := range vals {
		values = append(values, base+v)
	}
	return values[0]
}
