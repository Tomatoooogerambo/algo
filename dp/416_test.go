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
	for i := range dp {
		dp[i] = make([]bool, sum+1)
	}

	dp[0][0] = true

	for i, v := range nums {
		for j := 0; j <= sum; j++ {
			dp[i+1][j] = j >= v && dp[i][j-v] || dp[i][j]
		}
	}

	return dp[l][sum]
}
