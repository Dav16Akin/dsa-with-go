package datastructures

import "fmt"

func RunTwoDimentionalSlices() {
	// 2 DIMENTIONAL ARRAY
	var TwoDArray [8][8]int
	TwoDArray[3][6] = 18
	TwoDArray[7][4] = 3
	// fmt.Println(TwoDArray)


	// 2 DIMENTIONAL SLICE
	var rows int = 7
	var cols int = 9

	var twodslices = make([][]int, rows)
	for i := range twodslices {
		twodslices[i] = make([]int, cols)
	}

	// fmt.Println(twodslices)

	var arr = []int{5, 6, 7, 8, 9}
	var slic1 = arr[:3]
	fmt.Println("slice1", slic1)
	var slic2 = arr[1:5]
	fmt.Println("slice2", slic2)
	var slic3 = append(slic2, 12)
	fmt.Println("slice3", slic3)
}
