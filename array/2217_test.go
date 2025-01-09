package array

import (
	"math"
	"testing"
)

func TestKuhPalindrome(t *testing.T) {
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
	for i := 0; i < len(queries); i++ {
		ans[i] = findKthPalidrom(queries[i], intLength)
	}

	return ans
}

func findKthPalidrom(query, intL int) int64 {
	if intL == 1 {
		return int64(query)
	}

	if intL == 2 {
		return int64(query*10 + query)
	}

	N := int(0)
	if intL%2 == 1 {
		N = int((intL+1)/2 - 1)
	} else {
		N = int(intL/2 - 1)
	}

	index := query / int64(math.Pow10(N))
	nextQuery := query % int64(math.Pow10(N))
	return (index+1)*int64(math.Pow10(intL)) + (index + 1) + findKthPalidrom(nextQuery, intL-2)
}
