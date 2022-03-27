package main

import "fmt"

type Person struct {
	name string
	age  int
}

func CallByValueCheck() {
	person := Person{
		name: "sathish",
		age:  35,
	}

	modifyFails(person)
	fmt.Println(person)
}

func modifyFails(p Person) {
	p.name = "bob"
}
