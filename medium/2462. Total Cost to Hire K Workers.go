package medium

import (
	"container/heap"
)

type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MinHeap) Peek() int {
	return (*h)[0]
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func totalCost(costs []int, k int, candidates int) int64 {
	leftHeap := &MinHeap{}
	rightHeap := &MinHeap{}
	heap.Init(leftHeap)
	heap.Init(rightHeap)

	left, right := 0, len(costs)-1
	for i := 0; i < candidates && left <= right; i++ {
		heap.Push(leftHeap, costs[left])
		left++
		if left <= right {
			heap.Push(rightHeap, costs[right])
			right--
		}
	}

	var sum int64 = 0
	for i := 0; i < k; i++ {
		var minCost int
		if rightHeap.Len() == 0 || (leftHeap.Len() > 0 && leftHeap.Peek() <= rightHeap.Peek()) {
			minCost = heap.Pop(leftHeap).(int)
			if left <= right {
				heap.Push(leftHeap, costs[left])
				left++
			}
		} else {
			minCost = heap.Pop(rightHeap).(int)
			if left <= right {
				heap.Push(rightHeap, costs[right])
				right--
			}
		}
		sum += int64(minCost)
	}

	return sum
}

// Time: O((Candidates+K)*Log(candidates))
// Space: O(candidates)
