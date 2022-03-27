package main

import (
	"fmt"
)

func makeMult(base int) func(int) int {
	return func(factor int) int { return base * factor }
}

func FunctionReturnFunction() {
	base2 := makeMult(2)
	base3 := makeMult(3)

	fmt.Println("base2(5):", base2(5))
	fmt.Println("base3(10):", base3(10))
}
