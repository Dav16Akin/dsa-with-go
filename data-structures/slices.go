package datastructures

import "fmt"

func twiceValue(slice []int) {
	for i, value := range slice {
		slice[i] = 2 * value
	}
}

func RunSlice() {
	var slice = []int{1,3,5,6}
	slice = append(slice, 8)
	fmt.Println("capacity: ", cap(slice))
	fmt.Println("length: ", len(slice))
	twiceValue(slice)
	fmt.Println(slice)
}