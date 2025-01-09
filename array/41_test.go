package array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_FirstMissingPositive(t *testing.T) {
	a := []int{1, 2, 0}
	res := firstMissingPositive(a)
	assert.Equal(t, 3, res)

	a = []int{3, 4, -1, 1}
	res = firstMissingPositive(a)
	assert.Equal(t, 2, res)

	a = []int{1, 2, 3}
	res = firstMissingPositive(a)
	assert.Equal(t, 4, res)

	//[-1,4,2,1,9,10]
	a = []int{-1, 4, 2, 1, 9, 10}
	res = firstMissingPositive(a)
	assert.Equal(t, 3, res)

	a = []int{1, 1}
	res = firstMissingPositive(a)
	assert.Equal(t, 2, res)
}

// 使用hash表的方式完成
// 原地将数组作为一个哈希表， 数组的元素i 应该放在i-1的位置
// 3   4  -1  1
// -1  4  3   1
// -1  1  3   4
// -1  1  3   4
// 1  -1  3   4
func firstMissingPositive(nums []int) int {
	n := len(nums)
	for i := 0; i < n; i++ {
		for nums[i] > 0 && nums[i] < n && nums[i] != i+1 && nums[i] != nums[nums[i]-1] {
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
		}
	}

	for i, v := range nums {
		if i+1 != v {
			return i + 1
		}
	}

	return nums[n-1] + 1
}
