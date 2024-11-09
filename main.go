package main

import (
	"fmt"
)

func main() {
	name := "Hello World!"
	fmt.Printf("%s\n", name)
	createMatrix(10, 5)
}

// Создание матрицы
func createMatrix(rows, cols int) [][]int {
	matrix := make([][]int, rows)
	for i := 0; i < rows; i++ {
		matrix[i] = make([]int, cols)
		for j := 0; j < cols; j++ {
			matrix[i][j] = i * j
		}
	}

	for i := 0; len(matrix) > i; i++ {
		fmt.Println(matrix[i])
	}

	return matrix
}
