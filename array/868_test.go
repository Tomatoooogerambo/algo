package array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBinaryGap(t *testing.T) {
	n := 22
	res := binaryGap(n)
	assert.Equal(t, 2, res)
	n = 5
	res = binaryGap(n)
	assert.Equal(t, 2, res)
	n = 8
	res = binaryGap(n)
	assert.Equal(t, 0, res)
}

// 22
// 10110
// 0 11
// 1 5
// 1 2
// 0
// 1
func binaryGap(n int) int {
	max := 0
	arr := make([]int, 0)
	for n != 0 {
		arr = append(arr, n%2)
		n /= 2
	}
	for i := 0; i < len(arr); i++ {
		if arr[i] == 1 {
			for j := i + 1; j < len(arr); j++ {
				if arr[j] == 1 {
					if j-i > max {
						max = j - i
					}
					i = j - 1
					break
				}
			}
		}
	}
	return max
}
