package twodimension

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// [[1,4,7,11,15],[2,5,8,12,19],[3,6,9,16,22],[10,13,14,17,24],[18,21,23,26,30]]
func TestSearchMatrix(t *testing.T) {
	matrix := [][]int{{1, 4, 7, 11, 15}, {2, 5, 8, 12, 19}, {3, 6, 9, 16, 22}, {10, 13, 14, 17, 24}, {18, 21, 23, 26, 30}}
	res := searchMatrix(matrix, 5)
	assert.Equal(t, true, res)

	res = searchMatrix(matrix, 20)
	assert.Equal(t, false, res)
}

func searchMatrix(matrix [][]int, target int) bool {
	m := len(matrix)
	n := len(matrix[0])

	i, j := 0, n-1
	for i < m && j >= 0 {
		if matrix[i][j] == target {
			return true
		}

		if matrix[i][j] < target {
			i++
		} else {
			j--
		}
	}

	return false
}
