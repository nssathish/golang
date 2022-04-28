package main

import (
	"fmt"
	"time"
)

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

type Counter struct {
	total       int
	lastUpdated time.Time
}

func (p Person) ToString() string {
	return fmt.Sprintf("%s %s, Age %d", p.FirstName, p.LastName, p.Age)
}

func (c *Counter) Increment() {
	c.total++
	c.lastUpdated = time.Now()
}

func (c Counter) ToString() string {
	return fmt.Sprintf("%d %v", c.total, c.lastUpdated)
}

func doUpdateWrong(c Counter) {
	c.Increment()
	fmt.Println("doUpdateWrong()", c.ToString())
}

func doUpdateRight(c *Counter) {
	c.Increment()
	fmt.Println("doUpdateRight()", c.ToString())
}

type IntTree struct {
	val         int
	left, right *IntTree
}

func (it *IntTree) Insert(val int) *IntTree {
	if it == nil {
		return &IntTree{val: val}
	}
	if val < it.val {
		it.left = it.left.Insert(val)
	} else {
		it.right = it.right.Insert(val)
	}
	return it
}

func (it *IntTree) Contains(val int) bool {
	switch {
	case it == nil:
		return false
	case val < it.val:
		return it.left.Contains(val)
	case val > it.val:
		return it.right.Contains(val)
	default:
		return true
	}
}

type Adder struct {
	start int
}

func (a Adder) AddTo(value int) int {
	return a.start + value
}

func MethodsAreFunctions() {
	myAdder := Adder{start: 0}
	fmt.Println(myAdder.AddTo(10))

	methodAsVar := myAdder.AddTo
	fmt.Println(methodAsVar(5))

	//method expression
	F := Adder.AddTo
	fmt.Println(F(myAdder, 100))

	//This is type declaration
	//*** NOT INHERITANCE ***
	type Score int
	type HighScore Score

	var i int = 300
	var s Score = 100
	var hs HighScore = 200

	//hs = s  // - compilation error
	//s = i   // - compilation error
	s = Score(i)
	hs = HighScore(s)
	fmt.Println(hs)
}

func IotaForEnums() {
	type MailCategory int
	const (
		Uncategorized MailCategory = iota
		Personal
		Spam
		Social
		Advertisements
	)
	fmt.Println(Uncategorized)
	fmt.Println(Advertisements)

	//assign a literal expression to a constant
	type BitField int
	const (
		Field1 BitField = 1 << iota //assigned 1 -> 2^0
		Field2                      //assigned 2 -> 2^1
		Field3                      //assigned 4 -> 2^2
		Field4                      //assigned 8 -> 2^3
		//because "<<" left shifts the ON bit
	)
	fmt.Println(Field1)
	fmt.Println(Field4)
}

type Employee struct {
	Name string
	ID   int
}

func (e Employee) Description() string {
	return fmt.Sprintf("Name: %s and ID: %v", e.Name, e.ID)
}

type Manager struct {
	Employee
	Reports []Employee
}

type Inner struct {
	X int
}
type Outer struct {
	Inner
	X int
}

func EmbeddingForComposition() {
	manager := Manager{
		Employee: Employee{
			Name: "Seethu",
			ID:   1,
		},
		Reports: []Employee{},
	}
	fmt.Println(manager.ID)
	fmt.Println(manager.Description())

	outer := Outer{
		Inner: Inner{
			X: 20,
		},
		X: 10,
	}
	fmt.Println(outer.X)
	fmt.Println(outer.Inner.X)
}

type AInner struct {
	A int
}

func (i AInner) IntPrinter(value int) string {
	return fmt.Sprintf("Inner: %d", value)
}
func (i AInner) Double() int {
	fmt.Println("Inner double:")
	return i.A * 2
}

type AOuter struct {
	AInner
	S string
}

func (O AOuter) IntPrinter(value int) string {
	return fmt.Sprintf("Outer: %d", value)
}

func (O AOuter) Double() int {
	return O.A * 2
}
func EmbeddingIsNotInheritance() {
	o := AOuter{
		AInner: AInner{
			A: 10,
		},
		S: "Hello",
	}
	fmt.Println(o.IntPrinter(5))
	fmt.Println(o.Double())
}
