package main

import (
	"fmt"
	"math/rand"
)

var counter = 0 //package level variable

func allAboutBlocks() {
	fmt.Println("All about Blocks")

	x := 10
	if x > 5 {
		fmt.Println(x)
		x := 5
		fmt.Println(x)
	}
	fmt.Println(x)

	y, z := 5, 10
	if y <= 5 {
		y, z := 10, 20
		fmt.Println(y, z)
	}
	fmt.Println(y, z)

	//special scoped variables in if. .else block
	if randInt := rand.Intn(10); randInt < 5 {
		fmt.Println("Special scoped random integer: ", randInt)
	} else {
		fmt.Println("Special scoped random integer > 5: ", randInt)
	}
	//fmt.Println(randInt) //randInt is scoped only to the if. .else block see the ":="

	/*
			fmt := 100
			fmt.Println(fmt) // this is not possible
			display(fmt) // this is possible
		//shadow-linter helps to find the shadow variables
		//go install shadow from the golang repo
		//go vet ./... and shadow ./... - will give find such variables.

	*/

	//the universal block
	//Don't do this in the production code **
	fmt.Println(true)
	//true := 100 'true' is a variable in the "Universal Block" of Go
	//So, it can be given any value
	fmt.Println(true)

	for i := 0; i < 10; i++ {
		var n = rand.Intn(10)
		if n == 0 {
			fmt.Println("The value ", n, " is too low")
		} else if n < 5 {
			fmt.Println("The value ", n, " is average")
		} else {
			fmt.Println("The value ", n, " is high")
		}
	}

	//for loop with iterator index
	for _, i := range []int{1, 2, 3, 4, 5} {
		fmt.Printf("%v ", i)
	}
	fmt.Printf("\n")

	m := map[string][]string{
		"a": {"apple", "america", "andhra"},
		"b": {"ball", "bombay", "batman"},
	}

	//for range
	for key, value := range m {
		fmt.Println("Key: ", key, " and Value: ", value)
	}

	//for condition
	var i = 0
	for i < 10 {
		fmt.Println("i: ", i)
		i++
	}

	//for infinite
	for {
		i++
		fmt.Println("i = ", i)
		if i < 20 {
			continue
		} else {
			break
		}
	}

	for i := 0; i < 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("FizzBuzz")
			continue
		}

		if i%3 == 0 {
			fmt.Println("Buzz")
			continue
		}

		if i%5 == 0 {
			fmt.Println("Fizz")
			continue
		}
		fmt.Println(i)
	}

	for idx, val := range []int{10, 20, 30, 40, 50, 60} { // always, will take index and value
		//use "_" - underscore if you don't want to name the index of the array or slice
		fmt.Println("value at ", idx, " ", val)
	}

	//=======================for-range of strings is quite different=============
	var someValues = []string{"hello", "apple🤌_"}
	for _, value := range someValues {
		fmt.Println("string value: ", value)
		for idx, val := range value {
			fmt.Println(idx, val, string(val))
		}
	}
	//though the strings are made of bytes
	//for-range loops through the "runes" instead of the bytes
	//output will be as follows
	/*
		0 97 a
		1 112 p
		2 112 p
		3 108 l
		4 101 e
		5 129292 🤌 --+
		9 95 _      --+ 4 indices skipped because byte range is 256 indices but 129292 is more than that
		//Go, made adjustments to fit that because it uses runes
	*/

	//Labeling 'for' statements
	samples := []string{"hello", "apple_!"}
outer:
	for _, sample := range samples {
		for i, r := range sample {
			fmt.Println(i, r, string(r))
			if r == 'l' {
				continue outer
			}
		}
	}
	allAboutSwitch()
}

func allAboutSwitch() {
	var startingAlphabet = "z"
	switch startingAlphabet {
	case "a":
		fmt.Println(startingAlphabet, " for apple")
		break
	case "b":
		fmt.Println(startingAlphabet, " for ball")
		break
	default:
		fmt.Println("Not sure about '", startingAlphabet, "' alphabet")
		break
	}

	words := []string{"a", "cow", "smile", "gopher", "octopus", "anthropologist"}
	for _, word := range words {
		switch length := len(word); length {
		case 1, 2, 3, 4:
			fmt.Println(word, " is a short word!")
		case 5:
			fmt.Println(word, "is exactly the right length:", length)
		case 6, 7, 8, 9: //'empty case' - means no implementation for this case condition
		//write something here otherwise, anything that matches this condition will not have
		//anything to do.
		//[IMPORTANT] //**** Unlike other languages Go's 'case' statements don't fall through to the next case**
		default:
			fmt.Println(word, "is a long word!")
		}
	}

	for _, word := range words {
		switch size := len(word); size {
		case 1, 2, 3, 4:
			fmt.Println(word, "is a small word!")
		case 5:
			fmt.Println(word, "is exactly of length 5!")
		case 6, 7, 8, 9:
			fallthrough
		default:
			fmt.Println(word, "is a large word!")
		}
	}

	//Blank Switches
	salutations := []string{"hello", "hi", "hey"}
	for _, word := range salutations {
		switch wordLength := len(word); {
		case wordLength > 5:
			fmt.Println(word, "is large")
		case wordLength <= 3:
			fmt.Println(word, "is small")
		default:
			fmt.Println(word, "is of exact size")
		}
	}
}

//func display(value int) {
//	fmt.Println(value)
//}
