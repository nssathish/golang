package main

import (
	"errors"
	"fmt"
)

func CalcRemainderAndMod(numerator, denominator int) (int, int, error) {
	if denominator == 0 {
		return 0, 0, errors.New("(" + fmt.Sprintf("%d", numerator) + "/" + fmt.Sprintf("%d", denominator) + ") cannot divide by 0")
	}
	quotient := numerator / denominator
	mod := numerator % denominator
	return quotient, mod, nil

}
func BasicErrorHandling() {
	quotient, modulus, err := CalcRemainderAndMod(1, 0)
	fmt.Println("quotient:", quotient, "modulus:", modulus, "error:", err)
}
