package string

import (
	"fmt"
	"strings"
	"testing"
)

func TestWordSubsets(t *testing.T) {
	words1 := []string{"amazon", "apple", "facebook", "google", "leetcode"}
	words2 := []string{"e", "o"}

	ans := wordSubsets(words1, words2)
	fmt.Println(ans)

	words1 = []string{"amazon", "apple", "facebook", "google", "leetcode"}
	words2 = []string{"l", "e"}
	ans = wordSubsets(words1, words2)
	fmt.Println(ans)
}

// ["amazon","apple","facebook","google","leetcode"]  ["lo","eo"]
// ["amazon","apple","facebook","google","leetcode"], words2 = ["l","e"]
// words1 = ["amazon","apple","facebook","google","leetcode"], words2 = ["e","o"]
func wordSubsets(words1 []string, words2 []string) []string {
	ans := make([]string, 0)

	// 构建words2的keysmap
	keys := make(map[rune]int)
	for _, v := range words2 {
		freq := getFreq(v)
		for k, c := range freq {
			if keys[k] < c {
				keys[k] = c
			}
		}
	}

	for i := 0; i < len(words1); i++ {
		if ok := isSubSet(words1[i], keys); ok {
			ans = append(ans, words1[i])
		}
	}
	return ans
}

func isSubSet(s string, keys map[rune]int) bool {
	for k, v := range keys {
		if strings.Count(s, string(k)) < v {
			return false
		}
	}

	return true
}

func getFreq(s string) map[rune]int {
	m := make(map[rune]int)
	for _, v := range s {
		m[v] += 1
	}

	return m
}
