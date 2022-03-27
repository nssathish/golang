package main

import (
	"fmt"
	"sort"
)

func FunctionAsParam() {
	type Person struct {
		FirstName string
		LastName  string
		Age       int
	}

	people := []Person{
		{"optimus", "prime", 100},
		{"bumble", "bee", 25},
		{"iron", "hide", 75},
	}

	//sort by lastname
	sort.Slice(people, func(i, j int) bool {
		return people[i].LastName < people[j].LastName
	})
	fmt.Println(people)

	//sort by age
	sort.Slice(people, func(i, j int) bool {
		return people[i].Age < people[j].Age
	})
	fmt.Println(people)
}
