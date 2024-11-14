package mathfunc

import "fmt"

// Создание матрицы
func CreateMatrix(rows, cols int) [][]int {
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
