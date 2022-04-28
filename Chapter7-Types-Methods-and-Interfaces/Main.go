package main

import (
	"fmt"
)

func main() {
	fmt.Println("Chapter 7 - Types, Methods and Interfaces in GO")
	UserDefinedTypesOperations()
	p := Person{
		FirstName: "Sathish",
		LastName:  "Kumar",
		Age:       34,
	}
	fmt.Println(p.ToString())

	var c Counter
	fmt.Println(c.ToString())
	c.Increment()
	fmt.Println(c.ToString())

	var _counter Counter
	fmt.Println("in main:", _counter.ToString())
	//call by value
	doUpdateWrong(_counter)
	fmt.Println("in main:", _counter.ToString())
	//call by reference
	doUpdateRight(&_counter)
	fmt.Println("in main:", _counter.ToString())

	var it *IntTree
	it = it.Insert(5)
	it = it.Insert(3)
	it = it.Insert(10)
	it = it.Insert(2)
	fmt.Println(it.Contains(2))
	fmt.Println(it.Contains(12))

	MethodsAreFunctions()
	IotaForEnums()
	EmbeddingForComposition()
	EmbeddingIsNotInheritance()

	//Interfaces
	InterfacesFundamentals()
	DuckTyping()
	EmbeddingAndInterfaces()
	NilAndInterfaces()
	EmptyInterfaces()
	// TypeAssertions()
	// TypeSwitches()
	// BridgesToInterfaces()
	// DependencyInjections()
}
