package main

import (
	"fmt"
)

func main() {

	// Create matrix 2x2
	matrix := [][]int{{1, 2}, {3, 4}}
	result := float64(matrix[0][0])
	// Divide all elements
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			if i == 0 && j == 0 {
				continue
			}
			result /= float64(matrix[i][j])
		}
	}
	// Print Result
	fmt.Println("Result:", result)
}
