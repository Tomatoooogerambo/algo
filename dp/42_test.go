package dp

import "testing"

func TestTrap(t *testing.T) {
	tests := []struct {
		height []int
		want   int
	}{
		{[]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}, 6},
	}

	for _, test := range tests {
		if got := trap(test.height); got != test.want {
			t.Errorf("trap(%v) = %v", test.height, got)
		}
	}
}

func trap(height []int) int {
	l := len(height)
	left, right := 0, l-1
	leftMax, rightMax := 0, 0
	count := 0

	for left < right {
		leftMax = max(leftMax, height[left])
		rightMax = max(rightMax, height[right])

		if height[left] < height[right] {
			count += leftMax - height[left]
			left++
		} else {
			count += rightMax - height[right]
			right--
		}
	}

	return count
}

func max(a, b int) int {
	if a < b {
		return b
	} else {
		return a
	}
}
