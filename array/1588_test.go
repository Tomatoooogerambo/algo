package array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSumOddLengthSubarrays(t *testing.T) {
	arr := []int{1, 4, 2, 5, 3}
	res := sumOddLengthSubarrays(arr)
	assert.Equal(t, 58, res)
}

// 1 4 2 5 3
// 0 1 5 7 12 15 --> 7 11 10
//
//	--> 15
//
// 1 4 2 5 3 7 11 10 15  =
func sumOddLengthSubarrays(arr []int) int {
	// consrtuct prfix sum
	sum := 0
	n := len(arr)

	sufix := make([]int, 1)
	for i := 0; i < n; i++ {
		sum += arr[i]
		if i != 0 {
			arr[i] += arr[i-1]
		}
		sufix = append(sufix, sum)
	}

	for i := 3; i < n+1; i += 2 {
		for j := 0; j+i < n+1; j++ {
			sum += sufix[i+j] - sufix[j]
		}
	}

	return sum
}
