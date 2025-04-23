package dfs

import (
	"fmt"
	"testing"
)

func Test46(t *testing.T) {
	nums := []int{1, 2, 3}
	res := permute(nums)
	fmt.Println(res)
}

func permute(nums []int) (res [][]int) {
	l := len(nums)
	dict := make(map[int]bool)
	path := make([]int, 0)

	var dfs func(int)
	dfs = func(idx int) {
		if len(path) == l {
			temp := make([]int, l)
			copy(temp, path)
			res = append(res, temp)
			return
		}

		for i := 0; i < l; i++ {
			c := nums[i]
			if !dict[c] {
				path = append(path, c)
				dict[c] = true
				dfs(i + 1)
				// 回溯
				path = path[:len(path)-1]
				dict[c] = false
			}
		}
	}
	dfs(0)
	return res
}
