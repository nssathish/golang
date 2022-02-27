package main

import "fmt"

func slicesMain() {
	//slice is kind of an array without size (dynamic array)
	var slice1 = []int{1, 2, 3, 4}
	fmt.Println(slice1)

	var slice2 []int
	fmt.Println(len(slice2))

	//fmt.Println(slice1 == slice2) // compiler error, cannot compare 2 slices
	fmt.Println(slice1 != nil)
	fmt.Println(slice2 == nil)

	fmt.Println(len(slice1))
	fmt.Println(len(slice2))

	slice2 = append(slice2, 1)
	fmt.Println(slice2)
	slice2 = append(slice2, 2)
	slice2 = append(slice2, 5)
	slice2 = append(slice2, 8)
	fmt.Println(slice2)
	slice2 = append(slice2, 9, 8, 6, 0, 3)
	fmt.Println(slice2)

	//append
	slice1 = append(slice1, slice2...)
	fmt.Println(slice1)

	lenVsCapacity()
	allAboutMake()
	slicingSlices()
	slicingArrays()
	stringRuneAndByte()
	allAboutMaps()
	allAboutStruts()
}

func allAboutStruts() {
}

func allAboutMaps() {
	var nilMap map[string][]int = make(map[string][]int)
	fmt.Println("Nilmap: ", nilMap)
	nilMap["hello"] = []int{1, 2, 3}
	fmt.Printf("nilMap: %v\n", nilMap)

	regularMap := map[string][4]int{
		"even":  {2, 4, 6},
		"odd":   {1, 2, 3, 4},
		"prime": {2, 3, 5, 7},
	}

	fmt.Printf("regularMap: %v\n", regularMap)

	//get/set values in Go
	someKeyValue := map[string][]int{}

	someKeyValue["hello"] = []int{1, 2}
	someKeyValue["world"] = []int{3, 4, 5}
	someKeyValue["seethu"] = []int{}
	someKeyValue["seethu"] = append(someKeyValue["seethu"], 5, 7, 9)

	fmt.Println(someKeyValue)

	someKeyValue["hello"][1]++
	fmt.Println(someKeyValue)

	//comma "ok" idiom
	v, ok := someKeyValue["hello"]
	fmt.Println(v, ok)
	u, ok := someKeyValue["sathish"]
	fmt.Println(u, ok)
	_a, ok := someKeyValue["seethu"]
	fmt.Println(_a, ok)

	//fmt.Println(delete(someKeyValue, "hello")) //doesn't return anything (always)
	delete(someKeyValue, "hello")
	fmt.Println(someKeyValue)

	//maps as sets
	//Go doesn't have "set"
	intSet := map[int]bool{}
	var values []int = []int{1, 2, 3, 4, 5, 7, 8, 9, 9, 1, 1}

	//*****CHoice of "map" here is to avoid having duplicates in the set (an important property)**********
	for i := 0; i < len(values); i++ {
		val := values[i]
		intSet[val] = true
	}
	//or
	for _, v := range values {
		intSet[v] = true // i thought _ is for index of the array or slice (here) but its not :(
	}
	fmt.Println(intSet[5])
	fmt.Println(intSet[500])

	if !intSet[500] {
		fmt.Println("500 is not in the set")
	}

	//using struct {} for "set element validation"

}

func stringRuneAndByte() {
	var name string = "hello 🌞"
	fmt.Println("name: ", name, " len(name): ", len(name))
	fmt.Println("name[0]:", name[0]) // it's a byte
	var Ordinal byte = name[0]
	fmt.Println("Ordinal(name[0]): ", Ordinal)
	fmt.Println("string(name[0]): ", string(name[0]))
	//name[0] = "t" //not allowed: string is immutable in Go too!!
	var value = 65
	var stringifyValue = string(value) // converts given int value to ASCII
	fmt.Println("value: ", value, " string(value): ", stringifyValue)

	var someRune rune = 'a'
	var runeAsString string = string(someRune)
	var someByte byte = 'b'
	var byteAsString string = string(someByte)
	fmt.Println("runeAsString: ", runeAsString, "byteAsString: ", byteAsString)

	//converting string to slices of runes and bytes
	var someString string = "hello 🌞"
	sliceOfRunes := []rune(someString) //slice of runes that contains hello 🌞
	sliceOfBytes := []byte(someString) //slice of bytes that contains hello 🌞

	fmt.Println("sliceOfRunes: ", sliceOfRunes)
	fmt.Println("sliceOfBytes: ", sliceOfBytes)
}

func slicingArrays() {
	x := [4]int{5, 6, 7, 8}
	y := x[:2]
	z := x[2:]
	x[0] = 10 //shares same memory so, y and z will have an impact
	fmt.Println("Slicing Array x")
	fmt.Println("x: ", x)
	fmt.Println("y: ", y)
	fmt.Println("z: ", z)

	//slicing arrays to have independent memory
	a := make([]int, 4)
	xCopy := copy(a, x[:])
	fmt.Println(xCopy, x)
	x[0] = 20
	fmt.Println("xCopy (copy count): ", xCopy)
	fmt.Println("copied slice of x: ", a)
	fmt.Println("x: ", x)
}

func lenVsCapacity() {
	var x []int
	fmt.Println(x, len(x), cap(x))
	x = append(x, 10)
	fmt.Println(x, len(x), cap(x))
	x = append(x, 20)
	fmt.Println(x, len(x), cap(x))
	x = append(x, 30)
	fmt.Println(x, len(x), cap(x))
	x = append(x, 40)
	fmt.Println(x, len(x), cap(x))
	x = append(x, 50)
	fmt.Println(x, len(x), cap(x))
}

func allAboutMake() {
	newSlice := make([]int, 5)
	printSliceProperties(newSlice)

	newSlice1 := make([]int, 5)
	newSlice1 = append(newSlice1, 10)
	fmt.Println(newSlice1)

	newSlice2 := make([]int, 0, 10)
	printSliceProperties(newSlice2)

	newSlice2 = append(newSlice2, 5, 6, 7, 8)
	printSliceProperties(newSlice2)

	//newSlice3 := make([]int, 5, 0) //len larger than cap in make
	//printSliceProperties(newSlice3)
}

func slicingSlices() {
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println(slice[:3])
	fmt.Println(slice[1:])
	fmt.Println(slice[1:3])
	fmt.Println(slice[:])
	//fmt.Println(slice[1:100]) //out of range excepiton (or) 'panic' here
	x := slice[:3]
	y := slice[1:]
	x[0] = 10
	y[1] = 20
	fmt.Println("x = ", x, " y = ", y, " slice = ", slice)

	y1 := slice[:2]
	printSliceProperties(y1)
	y1 = append(y1, 30)
	printSliceProperties(y1)
	printSliceProperties(slice)

	x2 := make([]int, 0, 5)
	x2 = append(x2, 1, 2, 3, 4)
	y2 := x2[:2]
	z2 := x2[2:]
	fmt.Println(cap(x2), cap(y2), cap(z2))
	y2 = append(y2, 30, 40, 50)
	x2 = append(x2, 60)
	z2 = append(z2, 70)
	fmt.Println(x2)
	fmt.Println(y2)
	fmt.Println(z2)

	fullSliceExpression()
}

func fullSliceExpression() {
	x := []int{1, 2, 3, 4, 5}
	y := x[:2:2]
	printSliceProperties(y)
	y = append(y, 30, 40, 50)
	printSliceProperties(y)
	z := x[2:4:4]
	printSliceProperties(z)
	z = append(z, 60, 70)
	printSliceProperties(z)
	printSliceProperties(x)
}

func printSliceProperties(slice []int) {
	fmt.Println("New slice: ", slice)
	fmt.Println("Len of newSlice: ", len(slice), " capacity of newSlice: ", cap(slice))
}
