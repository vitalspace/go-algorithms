package main

import "fmt"

func main() {
	// Create matrix 2x2
	matrix := [][]int{{1, 2}, {3, 4}}
	result := matrix[0][0]

	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			if i == 0 && j == 0 {
				continue
			}
			result %= matrix[i][j]
		}
	}
	// Print Result
	fmt.Println("Result:", result)
}
