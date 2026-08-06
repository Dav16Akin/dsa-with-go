package datastructures

import "fmt"

func powerSeries (num int) (int , int) {
	return num * num , num * num * num
}

func RunTuples () {
	var square int;
	var cube int;

	square, cube = powerSeries(3)

	fmt.Println("Square :", square, "Cube :", cube)
}