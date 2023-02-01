package main

import (
	"fmt"
)

func main() {
	// Create matrix 2x2
	matrix := [][]int{{2, 5}, {7, 3}}
	sum := 0
	// Add up alements
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			sum += matrix[i][j]
		}
	}
	// Print Result
	fmt.Println("Result:", sum)
}
