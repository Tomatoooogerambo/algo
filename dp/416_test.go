package dp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test416(t *testing.T) {
	nums := []int{1, 5, 11, 5}
	res := canPartition(nums)
	assert.True(t, res)
}

func canPartition(nums []int) bool {
	sum := 0
	for _, v := range nums {
		sum += v
	}

	if sum%2 != 0 {
		return false
	}

	sum /= 2

	l := len(nums)

	// 初始化 dp
	dp := make([][]bool, l+1)
	for i := 0; i < l+1; i++ {
		dp[i] = make([]bool, sum+1)
	}

	dp[0][0] = true

	for i := 0; i < l; i++ {
		for j := 0; j <= sum; j++ {
			if j >= nums[i] {
				dp[i+1][j] = dp[i][j-nums[i]] || dp[i][j]
			}
		}
	}

	return dp[l][sum]
}
