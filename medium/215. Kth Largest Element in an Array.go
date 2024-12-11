package medium

import (
	"container/heap"
)

type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *MinHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func findKthLargest(nums []int, k int) int {
	h := &MinHeap{}
	heap.Init(h)
	for _, val := range nums {
		if h.Len() < k {
			heap.Push(h, val)
		} else {
			if val > (*h)[0] {
				heap.Push(h, val)
				heap.Pop(h)
			}
		}
	}

	return heap.Pop(h).(int)
}

// Time: O(NlogK)
// Space O(K)
