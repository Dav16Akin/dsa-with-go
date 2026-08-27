package algorithms

import "fmt"

//QUESTION 1
//For an array of 10 elements with a random set of integers,
// identify the maximum and minimum.
// Calculate the complexity of the algorithm.

func findMinMax(arr[10] int) (int, int) {
	max := arr[0] 
	min := arr[0]	

	for i:= 1; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
		} 

		if arr[i] < min {
			min = arr[i]
		}
	}

	return min,max
}

func RunAlgorithmExcersise() {
	var arr = [10]int{5,4,7,8,3,9,-2,4,-1,8}
	var min , max = findMinMax(arr)
	fmt.Printf("min: %d, max: %d\n", min, max)
}



