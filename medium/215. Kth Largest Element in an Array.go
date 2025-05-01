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

func findKthLargest(nums []int, k int) int {
	return quickSelect(nums, 0, len(nums)-1, k-1)
}

func quickSelect(nums []int, start, end int, k int) int {
	if end == start {
		return nums[start]
	}

	randomIndex := start + rand.Intn(end-start+1)
	pivot := nums[randomIndex]
	swap(nums, randomIndex, end)

	currentLeft := start
	for i := start; i <= end; i++ {
		if nums[i] > pivot {
			swap(nums, i, currentLeft)
			currentLeft++
		}
	}

	swap(nums, end, currentLeft)
	currentRight := end
	for i := end; i > currentLeft; i-- {
		if nums[i] < pivot {
			swap(nums, i, currentRight)
			currentRight--
		}
	}

	if currentLeft <= k && k <= currentRight {
		return pivot
	} else if currentLeft > k {
		return quickSelect(nums, start, currentLeft-1, k)
	} else {
		return quickSelect(nums, currentRight+1, end, k)
	}
}

func swap(nums []int, i, j int) {
	nums[i], nums[j] = nums[j], nums[i]
}

// Time: N + N/2 + N/4 + ... = O(N)
// Space: O(1)
