package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
)

//interface names end with "er"
type Stringer interface {
	String() string
}

type Logic interface {
	Process(data string) string
}
type LogicProvider struct {
}

func (lp LogicProvider) Process(data string) string {
	//businesslogic
	return ""
}

type Client struct {
	L Logic
}

func (c Client) Program() {
	// get data
	c.L.Process("")
}

func InterfacesFundamentals() {
	c := Client{
		L: LogicProvider{},
	}
	c.Program()
}

//storage handler interface
type FileStorageHandler interface {
	Save(content string) string
	Download(file string) string
}

//Blob storage handler
type BlobStorageHandler struct {
	message string
}

func (b BlobStorageHandler) Save(content string) string {
	return fmt.Sprintf("Saved Blob Content: %s", content)
}
func (b BlobStorageHandler) Download(file string) string {
	return fmt.Sprintf("Downloaded file Blob content: %s", file)
}
func (b BlobStorageHandler) Check(file string) string {
	return fmt.Sprintf("Check blob: %s", file)
}

//Disk storage handler
type DiskStorageHandler struct {
	message string
}

func (d DiskStorageHandler) Save(content string) string {
	return fmt.Sprintf("Sved content to disk: %s", content)
}
func (d DiskStorageHandler) Download(file string) string {
	return fmt.Sprintf("Downloaded file from disk: %s", file)
}

func SaveFile(h FileStorageHandler, content string) string {
	return h.Save(content)
}
func DownloadFile(h FileStorageHandler, file string) string {
	return h.Download(file)
}
func DuckTyping() {
	b := BlobStorageHandler{
		message: "some blob",
	}
	fmt.Println(SaveFile(b, "hello"))
	fmt.Println(DownloadFile(b, "hello"))

	d := DiskStorageHandler{
		message: "something in disk",
	}
	fmt.Println(SaveFile(d, "hai"))
	fmt.Println(DownloadFile(d, "hai"))
}

type BlobReader interface {
	ReadContent(path string) string
}
type AzureBlobReader struct{}

func (a AzureBlobReader) ReadContent(path string) string {
	return fmt.Sprintf("Read blob from path %s", path)
}

type FileReader interface {
	ReadContent(path string) string
}
type AzureFileReader struct{}

func (af AzureFileReader) ReadContent(path string) string {
	return fmt.Sprintf("Read file from path %s", path)
}

type GenericReader interface {
	BlobReader
	FileReader
}

func CloudRead(g GenericReader, path string) string { return g.ReadContent(path) }
func EmbeddingAndInterfaces() {
	b := AzureBlobReader{}
	fmt.Println(CloudRead(b, "/"))
	af := AzureFileReader{}
	fmt.Println(CloudRead(af, "/file/"))
}

func NilAndInterfaces() {
	var s *string
	fmt.Println(s == nil) //should return true: nil is 0 for pointer
	var i interface{}
	fmt.Println(i == nil) //same as above: nil is 0 for interface
	i = s
	//fmt.Println(i == nil) //false. Because The *string type (concrete type) is non-nil
}

func EmptyInterfaces() {
	var i interface{}
	i = 10
	i = "Hello"
	i = [5]int{1, 2, 3, 5, 5}
	i = struct {
		Firstname string
		Lastname  string
	}{"Seethu", "Sathish"}

	fmt.Println(i)

	// one set of braces for the interface{} type
	//the other to instantiate an instance of the map
	m := map[string]interface{}{}
	contents, err := ioutil.ReadFile("sample.json")

	if err != nil {
		fmt.Println(err)
	}

	json.Unmarshal(contents, &m)
	fmt.Println(m)
	fmt.Println(m["Key1"])
}

//****** GO LACKS GENERICS *********
//empty interface type to the rescue

type LinkedList struct {
	val  interface{}
	Next *LinkedList
}

func (ll *LinkedList) insert(pos int, value interface{}) *LinkedList {
	if ll == nil || pos == 0 {
		return &LinkedList{
			val:  value,
			Next: ll,
		}
	}
	ll.Next = ll.Next.insert(pos-1, value)
	return ll
}

func TypeAssertions() {
	type MyInt int

	var i interface{}
	var mine MyInt = 20
	i = mine

	i2 := i.(MyInt)
	fmt.Println(i2 + 1)

	if i3, ok := i.(string); ok { //run time error
		fmt.Println(i3)
	} else {
		fmt.Println(fmt.Errorf("unexpected type for %v", i))
	}

	if i4, ok := i.(int); ok { //run time error
		fmt.Println(i4 + 1)
	} else {
		fmt.Println(fmt.Errorf("unexpected type for %v", i))
	}
}

type MyInt int

func CheckType(i interface{}) string {
	//var MyInt int
	switch j := i.(type) {
	case nil:
		return "i is nil, type of j is interface{}"
	case int, MyInt, io.Reader, string, bool, rune:
		return fmt.Sprintf("j is of type %T\n", j)
	default:
		return fmt.Sprintln("no idea what i is, so j is of type interface{}")
	}
}

type treeNode struct {
	val    treeVal
	lChild *treeNode
	rChild *treeNode
}

//treeVal defines an "unexported marker interface"
//that makes it clear which types can be assigned
//to val in treeNode
type treeVal interface {
	isToken()
}

type number int

func (number) isToken() {}

type operator func(int, int) int

func (operator) isToken() {}
func (o operator) process(n1, n2 int) int {
	return o(n1, n2)
}

// tint := int(20) // throws compiler error: outside a function 'var' keyword is mandatory
// fmt.Println(tint)

var operators = map[string]operator{
	"+": func(n1, n2 int) int {
		return n1 + n2
	},
	"-": func(n1, n2 int) int {
		return n1 - n2
	},
	"*": func(n1, n2 int) int {
		return n1 * n2
	},
	"/": func(n1, n2 int) int {
		return n1 / n2
	},
}

func walkTree(tree *treeNode) (int, error) {
	switch val := tree.val.(type) {
	case nil:
		return 0, errors.New("invalid expression")
	case number:
		// we know that tree.val is of type number
		// so return the int value
		return int(val), nil
	case operator:
		// we know that tree.val is of type operator
		// so find the values of the left/right children
		// call process() on the values
		left, err := walkTree(tree.lChild)
		if err != nil {
			return 0, err
		}
		right, err := walkTree(tree.rChild)
		if err != nil {
			return 0, err
		}

		return val.process(left, right), nil
	default:
		// if a new treeVal type is defined but
		// walk tree wasn't updated
		// to process it, this detects it
		return 0, errors.New("uknown node type")
	}
}
func parse(s string) (*treeNode, error) {
	return &treeNode{
		val: operators["+"],
		lChild: &treeNode{
			val:    operators["*"],
			lChild: &treeNode{val: number(5)},
			rChild: &treeNode{val: number(10)},
		},
		rChild: &treeNode{val: number(20)},
	}, nil
}

func TypeSwitches() {
	var i interface{}
	var mine MyInt = 20
	i = mine
	fmt.Println(CheckType(i))

	parseTree, err := parse("5*10+20")
	if err != nil {
		panic(err)
	}
	result, err := walkTree(parseTree)
	fmt.Println(result, err)
}
