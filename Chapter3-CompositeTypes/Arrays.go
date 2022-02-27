package main

import (
	"fmt"
)

func arrayMain() {
	var numbers []int = []int{2, 3, 4, 5, 6}
	fmt.Println(numbers)

	// numbers = [3]int{3, 4, 5}

	//sparse array
	//specify index and associated number to the index
	var sparseArray [12]int = [12]int{1, 5: 4, 6, 10: 100, 15}

	fmt.Println(sparseArray)

	var dynamicArray = [...]int{8, 9, 91, 12}
	fmt.Println(dynamicArray)

	//array equality check
	var arr1 = [...]int{1, 2, 3}
	var arr2 = [3]int{1, 2, 3}

	fmt.Println(arr1 == arr2)

	//var size int = 10
	//var arr3 = [size]int{1: 2, 9: 5} //array size cannot be a varibale

	var size int = len(arr1)
	fmt.Println(size)
}
