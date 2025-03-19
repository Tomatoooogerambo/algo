package sort

import (
	"container/heap"
	"fmt"
	"testing"
)

type MaxHeap []int

func (max MaxHeap) Len() int {
	return len(max)
}

func (max MaxHeap) Less(i, j int) bool {
	return max[i] > max[j]
}

func (max MaxHeap) Swap(i, j int) {
	max[i], max[j] = max[j], max[i]
}

func (max *MaxHeap) Push(n any) {
	*max = append(*max, n.(int))
}

func (max *MaxHeap) Pop() any {
	old := *max
	n := len(old)
	x := old[n-1]
	*max = old[0 : n-1]
	return x
}

type MinHeap struct {
	MaxHeap
}

func (m MinHeap) Less(i, j int) bool {
	return m.MaxHeap[i] < m.MaxHeap[j]
}

func TestOriginHeap(t *testing.T) {
	max := &MaxHeap{}
	min := &MinHeap{}
	heap.Init(max)
	heap.Init(min)
	a := []int{7, 3, 5, 1, 4}
	// a := []int{3, 1, 6, 5, 4, 2}

	for i := 0; i < len(a); i++ {
		if max.Len() == 0 || a[i] < (*max)[0] {
			heap.Push(max, a[i])
		} else {
			heap.Push(min, a[i])
		}

		if diff(max.Len(), min.Len()) == 2 {
			if max.Len() > min.Len() {
				n := heap.Pop(max)
				heap.Push(min, n)
			} else {
				n := heap.Pop(min)
				heap.Push(max, n)
			}
		}
	}

	if diff(max.Len(), min.Len()) == 0 {
		num := heap.Pop(max).(int) + heap.Pop(min).(int)
		fmt.Println(float64(num) / 2)
	} else if max.Len() > min.Len() {
		fmt.Println(float64(heap.Pop(max).(int)))
	} else {
		fmt.Println(float64(heap.Pop(min).(int)))
	}

	// heap.Push(max, 3)
	// heap.Push(max, 5)
	// heap.Push(max, 2)
	// heap.Push(max, 6)
	// heap.Push(max, 1)
	//
	// heap.Push(min, 3)
	// heap.Push(min, 5)
	// heap.Push(min, 2)
	// heap.Push(min, 6)
	// heap.Push(min, 1)
	//
	for max.Len() > 0 {
		fmt.Print(heap.Pop(max))
	}

	for min.Len() > 0 {
		fmt.Print(heap.Pop(min))
	}
}

func diff(a, b int) int {
	if a > b {
		return a - b
	} else {
		return b - a
	}
}
