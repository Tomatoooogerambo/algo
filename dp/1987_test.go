package dp

import (
	"fmt"
	"math"
	"testing"
)

func Test1987(t *testing.T) {
	binary := "1100100010101010100100000111110001111001000010000010010111011"

	res := numberOfUniqueGoodSubsequences(binary)
	fmt.Println(res)
}

func numberOfUniqueGoodSubsequences(binary string) int {
	l := len(binary)
	dp0 := 0
	dp1 := 0
	has0 := false
	mod := math.Pow(10, 9) + 7
	for i := l - 1; i >= 0; i-- {
		if '1'-binary[i] == 0 {
			dp1 = int(math.Mod(float64(dp1+dp0+1), mod))
		} else {
			dp0 = int(math.Mod(float64(dp1+dp0+1), mod))
			has0 = true
		}
	}

	if has0 {
		return dp1 + 1
	}
	return dp1
}
