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

func modifyMap() {
	m := map[int]string{
		1: "one",
		2: "two",
	}
	modMap(m)
	fmt.Println(m)
}
func modMap(m map[int]string) {
	m[2] = "hello"
	m[3] = "goodbye"
	delete(m, 1)
}

func modifySlice() {
	s := []int{1, 2, 3}
	modSlice(s)
	fmt.Println(s)
}
func modSlice(s []int) {
	for idx, val := range s {
		s[idx] = val * 2
	}
	s = append(s, 10)
}
