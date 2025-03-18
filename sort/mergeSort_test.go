package sort

import (
	"fmt"
	"testing"
)

func TestMerge(t *testing.T) {
	a := []int{3, 1, 4, 5, 2, 6}
	res := MergeSort(a)
	fmt.Println(res)
}

func MergeSort(nums []int) []int {
	return move(nums, 0, len(nums)-1)
}

func move(nums []int, i, j int) []int {
	if i == j {
		return []int{nums[i]}
	}

	mid := (i + j) / 2

	left := move(nums, i, mid)
	right := move(nums, mid+1, j)
	return merge(left, right)
}

func merge(left, right []int) []int {
	nums := make([]int, 0)

	i, j := 0, 0
	for i < len(left) && j < len(right) {
		if left[i] > right[j] {
			nums = append(nums, left[i])
			i++
		} else {
			nums = append(nums, right[j])
			j++
		}
	}
	for i < len(left) {
		nums = append(nums, left[i])
		i++
	}
	for j < len(right) {
		nums = append(nums, right[j])
		j++
	}
	return nums
}
