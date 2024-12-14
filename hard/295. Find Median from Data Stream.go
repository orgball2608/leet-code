package hard

import "container/heap"

type MaxHeap []int

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

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
	*h = old[0 : n-1]
	return x
}

type MedianFinder struct {
	MinHeap *MinHeap
	MaxHeap *MaxHeap
}

func Constructor() MedianFinder {
	maxHeap := &MaxHeap{}
	minHeap := &MinHeap{}
	heap.Init(maxHeap)
	heap.Init(minHeap)
	return MedianFinder{
		MinHeap: minHeap,
		MaxHeap: maxHeap,
	}
}

func (this *MedianFinder) AddNum(num int) {
	if this.MaxHeap.Len() == 0 || (*this.MaxHeap)[0] >= num {
		heap.Push(this.MaxHeap, num)
	} else {
		heap.Push(this.MinHeap, num)
	}

	if this.MinHeap.Len()+1 < this.MaxHeap.Len() {
		heap.Push(this.MinHeap, heap.Pop(this.MaxHeap))
	} else if this.MinHeap.Len() > this.MaxHeap.Len() {
		heap.Push(this.MaxHeap, heap.Pop(this.MinHeap))
	}
}

// Time: O(LogN)
// Space: O(N)

func (this *MedianFinder) FindMedian() float64 {
	if this.MaxHeap.Len() > this.MinHeap.Len() {
		return float64((*this.MaxHeap)[0])
	}
	return float64((*this.MaxHeap)[0]+(*this.MinHeap)[0]) / float64(2)
}

// Time: O(1)
// Space: O(N)

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */
