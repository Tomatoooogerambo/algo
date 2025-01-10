package array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxKelements(t *testing.T) {
	nums := []int{1, 10, 3, 3, 3}
	res := maxKelements(nums, 3)
	assert.Equal(t, int64(17), res)
	nums = []int{1000000000}
	res = maxKelements(nums, 100000)
	assert.Equal(t, int64(1500099992), res)
}

type MaxHeap struct {
	data []int
}

func (mh *MaxHeap) heapfiUp(index int) {
	for index > 0 {
		parent := (index - 1) / 2
		if mh.data[parent] < mh.data[index] {
			mh.data[parent], mh.data[index] = mh.data[index], mh.data[parent]
			index = parent
		} else {
			break
		}
	}
}

func (mh *MaxHeap) heapFiDown(index int) {
	lastIndex := len(mh.data) - 1

	for {
		maxIndex := index
		leftChild := index*2 + 1
		rightChild := index*2 + 2

		if leftChild <= lastIndex && mh.data[leftChild] > mh.data[maxIndex] {
			maxIndex = leftChild
		}

		if rightChild <= lastIndex && mh.data[rightChild] > mh.data[maxIndex] {
			maxIndex = rightChild
		}

		// 如果当前index已经是最大的，停止
		if maxIndex == index {
			break
		}

		mh.data[maxIndex], mh.data[index] = mh.data[index], mh.data[maxIndex]
		index = maxIndex
	}
}

func (mh *MaxHeap) Insert(num int) {
	mh.data = append(mh.data, num)
	lastIndex := len(mh.data) - 1
	mh.heapfiUp(lastIndex)
}

func (mh *MaxHeap) Extract() int {
	n := len(mh.data)
	if n == 0 {
		return 0
	}

	res := mh.data[0]
	mh.data[0] = mh.data[n-1]
	mh.data = mh.data[:n-1]
	mh.heapFiDown(0)
	return res
}

func maxKelements(nums []int, k int) (ans int64) {
	heap := &MaxHeap{data: make([]int, 0)}

	// 将数组堆化
	for _, v := range nums {
		heap.Insert(v)
	}

	for i := 0; i < k; i++ {
		num := heap.Extract()
		ans += int64(num)

		num = (num + 2) / 3
		heap.Insert(num)
	}
	return ans
}
