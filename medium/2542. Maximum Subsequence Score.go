package medium

import (
	"container/heap"
	"sort"
)

type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func maxScore(nums1 []int, nums2 []int, k int) int64 {
	n := len(nums1)
	pairs := make([][2]int, n)
	for i := 0; i < n; i++ {
		pairs[i] = [2]int{nums1[i], nums2[i]}
	}

	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i][1] > pairs[j][1]
	})

	h := &MinHeap{}
	heap.Init(h)

	sum := 0
	maxScore := int64(0)

	for i := 0; i < n; i++ {
		num1 := pairs[i][0]
		num2 := pairs[i][1]

		heap.Push(h, num1)
		sum += num1

		if h.Len() > k {
			sum -= heap.Pop(h).(int)
		}

		if h.Len() == k {
			score := int64(sum) * int64(num2)
			if score > maxScore {
				maxScore = score
			}
		}
	}

	return maxScore
}
