package complexities

import "fmt"

func RunQuadratic() {
	//MULTIPLICATION TABLE
	var k , l int
	for k = 1; k <= 10; k++ {
		fmt.Println("Multiplication Table", k)
		for l = 1; l <= 10; l++ {
			println(l * k)
		}
	}
}
