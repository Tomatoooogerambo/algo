package sort

import (
	"fmt"
	"testing"
)

func TestHeap(t *testing.T) {
	a := []int{3, 1, 4, 5, 6, 2}
	h := &MHeap{}
	h.list = make([]int, 0)
	for i := 0; i < len(a); i++ {
		h.add(a[i])
	}

	for h.size() > 0 {
		fmt.Print(h.pop())
	}
}

// 构造堆
type MHeap struct {
	list []int
}

func (h *MHeap) left(i int) int {
	return 2*i + 1
}

func (h *MHeap) right(i int) int {
	return 2*i + 2
}

func (h *MHeap) parent(i int) int {
	return (i - 1) / 2
}

func (h *MHeap) size() int {
	return len(h.list)
}

func (h *MHeap) swap(i, j int) {
	h.list[i], h.list[j] = h.list[j], h.list[i]
}

func (h *MHeap) shiftUp(i int) {
	for {
		p := h.parent(i)
		if p < 0 || h.list[i] <= h.list[p] {
			break
		}
		h.swap(i, p)
		i = p
	}
}

func (h *MHeap) shiftDown(i int) {
	for {
		l, r := h.left(i), h.right(i)
		cur := i
		if l < h.size() && h.list[l] > h.list[cur] {
			cur = l
		}
		if r < h.size() && h.list[r] > h.list[cur] {
			cur = r
		}

		// 如果一同操作下来 cur没有变化,说明已经到了推出的位置
		if cur == i {
			break
		}

		h.swap(i, cur)
		i = cur
	}
}

func (h *MHeap) add(num int) {
	h.list = append(h.list, num)
	h.shiftUp(h.size() - 1)
}

func (h *MHeap) pop() int {
	p := h.list[0]
	h.swap(0, h.size()-1)
	h.list = h.list[:h.size()-1]
	h.shiftDown(0)
	return p
}
