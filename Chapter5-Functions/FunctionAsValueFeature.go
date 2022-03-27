package main

import (
	"fmt"
	"strconv"
)

func functionsAreValues() {
	oMap := map[string]func(int, int) int{
		"+": add,
		"-": sub,
		"*": mul,
		"/": div,
	}
	//or
	/*
		type opFunctionType func(int, int) int
		oMap := map[string]opFunctionType {
			"+": add,
			"-": sub,
			"*": mul,
			"/": div,
		}
	*/
	expressions := [][]string{
		{"2", "+", "3"},
		{"2", "-", "3"},
		{"2", "*", "3"},
		{"2", "/", "3"},
		{"two", "+", "three"},
		{"5"},
	}

	for _, expression := range expressions {
		if len(expression) != 3 {
			fmt.Println("invalid expression: ", expression)
			continue
		}
		param1, err := strconv.Atoi(expression[0])

		if err != nil {
			fmt.Println(err)
			continue
		}
		operator := expression[1]
		opFunc, ok := oMap[operator]

		if !ok {
			fmt.Println("unsupported operator: ", operator)
			continue
		}

		param2, err := strconv.Atoi(expression[2])
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println(opFunc(param1, param2))
	}
}

func add(i int, j int) int { return i + j }
func sub(i int, j int) int { return i - j }
func mul(i int, j int) int { return i * j }
func div(i int, j int) int { return i / j }
