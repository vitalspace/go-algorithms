package main

import (
	"fmt"
)

func main() {
	matrix := [][]int{{5, 8}, {7, 9}}
	result := 1

	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			result *= matrix[i][j]
		}
	}
	fmt.Println("Result:", result)
}
