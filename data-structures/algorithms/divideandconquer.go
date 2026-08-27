package algorithms

import "fmt"

func fibonacci(k int) int {
	if k <= 1 {
		return 1
	}

	return fibonacci(k-1) + fibonacci(k-2)
}

func RunDivideAndConquer() {
	var m int
	for m = 0; m < 8 ; m++ {
		var fib = fibonacci(m)
		fmt.Println(fib)
	}
}