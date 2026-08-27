package datastructures

import "fmt"

func RunArrays() {
	// array declaration
	var arr = [5]int{1, 2, 4, 5, 6}


	// Traversing arrays
	var i int
	for i = 0; i < len(arr); i++ {
		fmt.Println("printing elements ", arr[i])
	}

	var value int
	for i, value = range arr {
		fmt.Println(" range ",value)
	}
}
