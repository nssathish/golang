package main

import (
	"fmt"
)

type person struct {
	name   string
	age    int
	weight float32
}

func main() {
	fmt.Println("3 / 5 = ", division(3, 5))
	guy := person{
		name:   "Sathish",
		age:    35,
		weight: 75.1,
	}
	someOne(guy)
	numerator, denominator, err := multipleReturnValues()
	fmt.Println(numerator, denominator, err)
	fmt.Println(divAndReminder(5, 3))
	functionsAreValues()
	fmt.Println(multiplyBy(2, []int{1, 2, 3, 4, 5}))
	fmt.Println("Function as parameters")
	FunctionAsParam()
	FunctionReturnFunction()
}
