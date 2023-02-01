package main

import (
	"fmt"
)

// create a matrix with 2 row and 3 colums
// Transpose

func main() {

	// Matrix 2 x 3
	a := [][]float64{{2, 3, 5}, {2, 6, 8}}
	fmt.Println("a =")
	for i := 0; i < len(a); i++ {
		fmt.Println(a[i])
	}
	fmt.Println("")

	// Transpose Matrix
	at := make([][]float64, len(a[0]))
	for i := 0; i < len(a[0]); i++ {
		at[i] = make([]float64, len(a))
		for j := 0; j < len(a); j++ {
			at[i][j] = a[j][i]
		}
	}

	fmt.Println("at =")
	for i := 0; i < len(at); i++ {
		fmt.Println(at[i])
	}

}
