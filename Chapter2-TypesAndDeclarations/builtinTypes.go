package main

import (
	"fmt"
	"math"
)

func main() {
	var flag bool
	fmt.Printf("default value of flag %v\n", flag)
	newFlag := true
	fmt.Println(newFlag)

	// var a int = 1000
	// var b int16 = 300

	//var c int32 = a + b

	var a int = 1
	var b int = 0
	// var c int = a / b

	// fmt.Println(c)

	c := a * b
	fmt.Println(c)

	b *= a
	c = b

	fmt.Println(c)

	fmt.Println(a >= b)

	/*
		Bit operation
		left shift <<
		right shift >>
	*/

	a >>= 1
	fmt.Println(a)
	a = 1
	a <<= 1
	fmt.Println(a)

	var d byte = 255 // 1111 1111
	d <<= 1          //254
	fmt.Println(d)
	d = 255
	d >>= 1 //127
	fmt.Println(d)

	fmt.Println(math.MinInt)

	ConstFunc()
}

const x int = 10
const (
	idKey   = "id"
	nameKey = "name"
)
const z = 20 * 10

func ConstFunc() {
	const y = "hello"

	fmt.Println(x)
	fmt.Println(y)

	// x += 1

	// y = "bye"

	// fmt.Println(x)
	// fmt.Println(y)

	const x1 = 10 // untyped - means data type is not mentioned
	var x2 int = x1
	var y2 float64 = x1
	var z2 byte = x1

	fmt.Println(x2, y2, z2)

	var items [5]int
	for i := 0; i < len(items); i++ {
		fmt.Println(items[i])
	}
}
