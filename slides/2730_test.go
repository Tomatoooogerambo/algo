package slides

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindLong(t *testing.T) {
	a := "52233"
	res := longestSemiRepetitiveSubstring(a)
	fmt.Println(res)
	assert.Equal(t, 4, res)

	a = "5494"
	res = longestSemiRepetitiveSubstring(a)
	fmt.Println(res)
	assert.Equal(t, 4, res)

	a = "111111"
	res = longestSemiRepetitiveSubstring(a)
	fmt.Println(res)
	assert.Equal(t, 2, res)

	a = "0001"
	res = longestSemiRepetitiveSubstring(a)
	fmt.Println(res)
	assert.Equal(t, 3, res)

	a = "4411794"
	res = longestSemiRepetitiveSubstring(a)
	fmt.Println(res)
	assert.Equal(t, 6, res)

	a = "524446"
	res = longestSemiRepetitiveSubstring(a)
	fmt.Println(res)
	assert.Equal(t, 4, res)
}

func longestSemiRepetitiveSubstring(s string) int {
	n := len(s)
	if n == 1 {
		return 1
	}

	count := 0
	left := 0
	lastRepeatIndex := -1 // 用于记录最近一次重复字符的位置

	for right := 1; right < n; right++ {
		// 如果当前字符和前一个字符重复
		if s[right] == s[right-1] {
			// 如果已经有一个重复字符，调整左边界
			if lastRepeatIndex != -1 {
				left = lastRepeatIndex + 1
			}
			// 更新最近一次重复字符的位置
			lastRepeatIndex = right - 1
		}

		// 更新最长子串的长度
		count = max(count, right-left+1)
	}

	return count
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
