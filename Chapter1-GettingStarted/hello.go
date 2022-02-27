package main

import "fmt"

func main() {
	fmt.Println("hello world!")
	fmt.Println("Doing the first calculation")
	var a = 10
	var b = 20

	fmt.Printf("a = %v\nb = %v\na + b = %v\n", a, b, a+b)

	fmt.Println("Hello, \n \"Sathish\"")

	var flag bool
	fmt.Printf("default value of flag %v\n", flag)
	newFlag := true
	fmt.Printf("value of newFlag is %v\n", newFlag)
	//newFlag := 10 //This implise that Go has compiler
	fmt.Println(newFlag)

}
