package slides

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// │ 输入：nums = [1,3,-1,-3,5,3,6,7], k = 3
// │ 输出：[3,3,5,5,6,7]
// │ 解释：
// │ 滑动窗口的位置                最大值
// │ ---------------               -----
// │ [1  3  -1] -3  5  3  6  7       3
// │  1 [3  -1  -3] 5  3  6  7       3
// │  1  3 [-1  -3  5] 3  6  7       5
// │  1  3  -1 [-3  5  3] 6  7       5
// │  1  3  -1  -3 [5  3  6] 7       6
// │  1  3  -1  -3  5 [3  6  7]      7
//

func TestMaxSlidingWindow(t *testing.T) {
	nums := []int{1, 3, -1, -3, 5, 3, 6, 7}
	res := maxSlidingWindow(nums, 3)
	assert.Equal(t, []int{3, 3, 5, 5, 6, 7}, res)

	nums = []int{1, -1}
	res = maxSlidingWindow(nums, 1)
	assert.Equal(t, []int{1, -1}, res)
}

func maxSlidingWindow(nums []int, k int) []int {
	res := []int{}
	deque := []int{}

	// 维护deque单调栈性质
	for i := 0; i < len(nums); i++ {
		for len(deque) > 0 && nums[deque[len(deque)-1]] <= nums[i] {
			deque = deque[:len(deque)-1]
		}

		deque = append(deque, i)

		if i-deque[0] >= k {
			deque = deque[1:]
		}
		if i >= k-1 {
			res = append(res, nums[deque[0]])
		}
	}

	return res
}
