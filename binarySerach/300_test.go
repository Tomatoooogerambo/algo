package binaryserach

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test300(t *testing.T) {
	nums := []int{10, 9, 2, 5, 3, 7, 101, 18}
	res := lengthOfLIS(nums)
	assert.Equal(t, 4, res)

	nums = []int{0, 1, 0, 3, 2, 3}
	res = lengthOfLIS(nums)
	assert.Equal(t, 4, res)
	nums = []int{7, 7, 7, 7, 7, 7, 7}
	res = lengthOfLIS(nums)
	assert.Equal(t, 1, res)
}

func lengthOfLIS(nums []int) int {
	g := make([]int, 0)
	for _, num := range nums {
		idx := sort.SearchInts(g, num)
		if idx == len(g) {
			g = append(g, num)
		} else {
			g[idx] = num
		}
	}

	return len(g)
}
