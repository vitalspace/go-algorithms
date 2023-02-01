package main

import (
	"fmt"
)

func main() {
	// Create Matrix 2x2
	matrix := [][]float64{{1, 2}, {3, 4}}
	result := matrix[0][0]
	// Subtract all elements
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			if i == 0 && j == 0 {
				continue
			}
			result -= matrix[i][j]
		}
	}
	// Print result
	fmt.Println("Result", result)
}
