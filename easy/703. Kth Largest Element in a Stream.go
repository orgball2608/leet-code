package easy

import "container/heap"

type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] } // Min-Heap
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

type KthLargest struct {
	heap *MinHeap
	k    int
}

func Constructor(k int, nums []int) KthLargest {
	h := &MinHeap{}
	for i := range nums {
		*h = append(*h, nums[i])
	}
	heap.Init(h)

	n := len(nums)
	for n > k {
		heap.Pop(h)
		n--
	}

	return KthLargest{
		heap: h,
		k:    k,
	}
}

func (this *KthLargest) Add(val int) int {
	heap.Push(this.heap, val)
	if this.heap.Len() > this.k {
		heap.Pop(this.heap)
	}

	return (*this.heap)[0]
}

/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */
