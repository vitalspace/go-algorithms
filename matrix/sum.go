package main

import (
	"fmt"
)

func main() {

	matrix := [][]int{{2, 5}, {7, 3}}
	sum := 0

	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			sum += matrix[i][j]
		}
	}

	fmt.Println("Result:", sum)

}
