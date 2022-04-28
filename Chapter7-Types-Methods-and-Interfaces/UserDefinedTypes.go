package main

import (
	"fmt"
	"strconv"
)

func UserDefinedTypesOperations() {
	type Person struct {
		firstName string
		lastName  string
		age       int
	}

	p := Person{
		firstName: "seethu",
		lastName:  "sathish",
		age:       33,
	}
	fmt.Println(p)

	type score int
	type Converter func(string) score
	var converter Converter = func(s string) score {
		value, _ := strconv.Atoi(s)
		return score(value)
	}

	fmt.Println(converter("100"))

	type TeamScores map[string]score
	var teamScores TeamScores = TeamScores{
		"seethu":  100,
		"sathish": 100,
	}
	fmt.Println(teamScores)
}
