package twodimension

import "testing"

func TestRotate(t *testing.T) {
	// init a 2 dimension array
	matrix := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	rotate(matrix)
	PrintMatrix(matrix)

	matrix2 := [][]int{{5, 1, 9, 11}, {2, 4, 8, 10}, {13, 3, 6, 7}, {15, 14, 12, 16}}
	PrintMatrix(matrix2)
	rotate(matrix2)
	PrintMatrix(matrix2)
}

// 2维数组双指针
func rotate(matrix [][]int) {
	l := len(matrix) // Length of the matrix (NxN)

	// Iterate over layers
	for rI := 0; rI < l/2; rI++ {
		for cI := rI; cI < l-1-rI; cI++ {
			// Swap the four corners of the layer
			matrix[rI][cI], matrix[l-1-cI][rI], matrix[l-1-rI][l-1-cI], matrix[cI][l-1-rI] = matrix[l-1-cI][rI], matrix[l-1-rI][l-1-cI], matrix[cI][l-1-rI], matrix[rI][cI]
		}
	}
}

func PrintMatrix(matrix [][]int) {
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			print(matrix[i][j], "  ")
		}
		println()
	}
	println("-----------")
}
