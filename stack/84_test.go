package stack

import (
	"fmt"
	"testing"
)

func TestArea(t *testing.T) {
	heights := []int{2, 1, 2}
	res := largestRectangleArea(heights)
	fmt.Println(res)
}

func largestRectangleArea(heights []int) (res int) {
	// 单调栈 维护heights中的
	heights = append([]int{0}, heights...)
	heights = append(heights, 0)
	stack := make([]int, 0)

	n := len(heights)

	for i := 0; i < n; i++ {
		// 推出条件
		for len(stack) > 0 && heights[stack[len(stack)-1]] > heights[i] {
			res = max(res, heights[stack[len(stack)-1]]*(i-stack[len(stack)-2]-1))
			stack = stack[:len(stack)-1]
		}

		stack = append(stack, i)
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
