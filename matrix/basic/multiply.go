package main

import (
	"fmt"
)

func main() {
	// Create Matrix 2x2
	matrix := [][]int{{5, 8}, {7, 9}}
	result := 1
	// Multiply all elements
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			result *= matrix[i][j]
		}
	}
	// Print Result
	fmt.Println("Result:", result)
}
