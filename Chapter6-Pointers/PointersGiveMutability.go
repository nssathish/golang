package main

import "fmt"

func MutabilityCheck() {
	var x *int32
	failedUpdate(x)
	fmt.Println(x)
}

func failedUpdate(a *int32) {
	var y int32 = 10
	//fmt.Println("inside failedUpdate(a):", *a) 'nil' value reference error (no memory allocated)
	a = &y
}

func MutabilityCheckSuccess() {
	var a = 10
	var b *int = &a
	failingUpdate(b)
	fmt.Println(*b)
	successfulUpdate(b)
	fmt.Println(*b)
}

func successfulUpdate(b *int) {
	*b = 50
}

func failingUpdate(c *int) {
	x := 20
	c = &x
}
