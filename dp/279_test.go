package dp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test279(t *testing.T) {
	target := 13
	res := numSquares(target)
	assert.Equal(t, 2, res)

	target = 12
	res = numSquares(target)
	assert.Equal(t, 3, res)

	target = 4
	res = numSquares(target)
	assert.Equal(t, 1, res)
}

func numSquares(n int) int {
	// gouzao candidate array
	nums := make([]int, 0)
	for i := 0; i*i <= n; i++ {
		nums = append(nums, i*i)
	}

	dp := make([]int, n+1)
	dp[0] = 0
	dp[1] = 1

	for i := 2; i < n+1; i++ {
		dp[i] = dp[i-1] + 1
		for j := 0; j < len(nums); j++ {
			if i < nums[j] {
				break
			}
			dp[i] = min(dp[i], dp[i-nums[j]]+1)
		}
	}

	return dp[n]
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
