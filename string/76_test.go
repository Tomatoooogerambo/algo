package string

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinWindow(t *testing.T) {
	s := "ADOBECODEBANC"
	tt := "ABC"

	res := minWindow(s, tt)
	assert.Equal(t, "BANC", res)

	// add case s="aa", tt="aa"
	s = "aa"
	tt = "aa"
	res = minWindow(s, tt)
	assert.Equal(t, "aa", res)
}

func minWindow(s string, t string) string {
	tb := make(map[rune]int)
	hb := make(map[rune]int)
	res := ""
	count := 0
	for _, v := range t {
		tb[v]++
	}

	for i, j := 0, 0; j < len(s); j++ {
		if _, ok := tb[rune(s[j])]; ok {
			hb[rune(s[j])]++
			if tb[rune(s[j])] >= hb[rune(s[j])] {
				count++

				for count == len(t) {
					if res == "" || len(res) > j-i+1 {
						res = s[i : j+1]
					}
					if _, ok := tb[rune(s[i])]; ok {
						hb[rune(s[i])]--
						if tb[rune(s[i])] > hb[rune(s[i])] {
							count--
						}
					}
					i++
				}
			}
		}
	}

	return res
}
