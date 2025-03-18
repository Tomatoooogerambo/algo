package sort

import (
	"fmt"
	"testing"
)

func TestHeap(t *testing.T) {
	a := []int{3, 1, 4, 5, 6, 2}
	h := &heap{}
	h.list = make([]int, 0)
	for i := 0; i < len(a); i++ {
		h.add(a[i])
	}

	for h.size() > 0 {
		fmt.Print(h.pop())
	}
}

type heap struct {
	list []int
}

func (h *heap) left(i int) int {
	return 2*i + 1
}

func (h *heap) right(i int) int {
	return 2*i + 2
}

func (h *heap) parent(i int) int {
	return (i - 1) / 2
}

func (h *heap) swap(i, j int) {
	h.list[i], h.list[j] = h.list[j], h.list[i]
}

func (h *heap) size() int {
	return len(h.list)
}

func (h *heap) shiftUp(i int) {
	for {
		p := h.parent(i)
		if p < 0 || h.list[p] >= h.list[i] {
			break
		}

		h.swap(p, i)
		i = p
	}
}

func (h *heap) shiftDown(i int) {
	for {
		l, r, min := h.left(i), h.right(i), i
		if l < h.size() && h.list[l] > h.list[min] {
			min = l
		}

		if r < h.size() && h.list[r] > h.list[min] {
			min = r
		}

		if i == min {
			break
		}

		h.swap(min, i)
		i = min
	}
}

func (h *heap) add(n int) {
	h.list = append(h.list, n)
	h.shiftUp(h.size() - 1)
}

func (h *heap) pop() int {
	max := h.list[0]
	h.swap(0, h.size()-1)
	h.list = h.list[:h.size()-1]
	h.shiftDown(0)
	return max
}
