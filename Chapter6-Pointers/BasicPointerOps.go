package main

import "fmt"

func BasicOps() {
	var x int32 = 10
	var y bool = false

	addressOfX := &x
	addressOfY := &y

	pointerToX := *(addressOfX)
	pointerToY := *(addressOfY)

	fmt.Println("address(x)", addressOfX)
	fmt.Println()
	fmt.Println("address(y)", addressOfX)
	fmt.Println()
	fmt.Println("pointer(x)", pointerToX)
	fmt.Println()
	fmt.Println("pointer(y)", pointerToY)
}

func DeclareAsPointerThrowingError() {
	var a *int

	//check for nullability
	if a != nil {
		fmt.Println(*a) //*a is nil this will throw panic
	}

	//*a = 10
	//fmt.Println(*a)
}

func DeclareNewPointer() {
	var b = new(int)

	if b != nil {
		fmt.Println(*b)
	} else {
		fmt.Println(nil)
	}
}

func InitializePointersInStruct() {
	type person struct {
		firstName  string
		secondName string
		middleName *string
	}

	newPerson := person{
		firstName:  "sathish",
		secondName: "kumar",
		//middleName: "subramanian", //this will throw compiler error
		middleName: stringp("subramanian"),
	}

	fmt.Println(newPerson)
	fmt.Println(*newPerson.middleName)
}

func stringp(s string) *string {
	return &s
}
