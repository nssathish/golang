package main

import (
	"archive/zip"
	"bytes"
	"consterr"
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

func QuotientRemainderDivisorByErrorsNewFunc(numerator, denominator int) (int, int, int, error) {
	if denominator == 0 {
		return 0, 0, 0, errors.New("cannot divide by 0")
	}
	quotient := numerator / denominator
	remainder := numerator % denominator

	return quotient, remainder, denominator, nil
}

func QuotientRemainderDivisorByErrorf(numerator, denominator int) (int, int, int, error) {
	if denominator == 0 {
		return 0, 0, 0, fmt.Errorf("cannot divide by 0")
	}
	quotient := numerator / denominator
	remainder := numerator % denominator

	return quotient, remainder, denominator, nil
}

func UsingStringsForSimpleErrors() {
	fmt.Println(QuotientRemainderDivisorByErrorf(20, 0))
	fmt.Println(QuotientRemainderDivisorByErrorf(3, 2))
	fmt.Println(QuotientRemainderDivisorByErrorsNewFunc(20, 0))
	fmt.Println(QuotientRemainderDivisorByErrorsNewFunc(5, 4))
}

// to signal the processing cannot continue due to a problem
// of the **** CURRENT STATE **** (basically a state management)
func SentinelErrors() {
	data := []byte("This is not a zip file")
	notAZipFile := bytes.NewReader(data)
	_, err := zip.NewReader(notAZipFile, int64(len(data)))
	if err == zip.ErrFormat { // " 'Err'Format is the sentinel error "
		fmt.Println(fmt.Errorf("zip file erorr"))
	}

	const (
		ErrFoo = consterr.Sentinel("foo error")
		ErrBar = consterr.Sentinel("bar error")
	)
}
