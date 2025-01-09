package array

import (
	"fmt"
	"math"
	"testing"
)

func TestKuhPalindrome(t *testing.T) {
	queryies := []int{1, 2, 3, 4, 5, 90}
	intLength := 3
	res := kthPalindrome(queryies, intLength)
	fmt.Println(res)
}

// 1 1 2 3 4 5 6 7 8 9                            1 * 9
// 2 11 22 33 44 55 66 77 88 99                   1 * 9
// 3 1x1 2x2 3x3 4x4 5x5 6x6 7x7 8x8 9x9          10* 9
// 4 1xx1 2xx2 3xx3 4xx4 5xx5 6xx6 7xx7 8xx8 9xx9 10* 9
// 5 1xxx1 2xxx2 3xxx3 4xxx4 5xxx5 6xxx6 7xxx7 8xxx8 9xxx9 100*9  query 70  5
//
//	query 30  3
func kthPalindrome(queries []int, intLength int) []int64 {
	ans := make([]int64, len(queries))
	l := intLength / 2
	if intLength%2 == 1 {
		l = intLength/2 + 1
	}

	start := int(math.Pow10(l-1) - 1)
	limit := int(math.Pow10(l) - 1)
	for i := 0; i < len(queries); i++ {
		if start+queries[i] > int(limit) {
			ans[i] = int64(-1)
		} else {
			ans[i] = int64(makePalindrome(start+queries[i], intLength))
		}
	}

	return ans
}

func makePalindrome(query, intLength int) int {
	// 根据intLength奇偶返回
	temp := query
	if intLength%2 == 1 {
		temp /= 10
	}

	for temp > 0 {
		query = query*10 + temp%10
		temp = temp / 10
	}

	return query
}
