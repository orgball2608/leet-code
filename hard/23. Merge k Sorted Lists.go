package hard

import (
	"container/heap"
)

type MinHeap []*ListNode

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i].Val < h[j].Val }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(*ListNode))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func mergeKLists(lists []*ListNode) *ListNode {
	minHeap := &MinHeap{}
	heap.Init(minHeap)

	for _, list := range lists {
		if list != nil {
			heap.Push(minHeap, list)
		}
	}

	dummy := &ListNode{}
	current := dummy

	for minHeap.Len() > 0 {
		node := heap.Pop(minHeap).(*ListNode)
		current.Next = node
		current = current.Next
		if node.Next != nil {
			heap.Push(minHeap, node.Next)
		}
	}

	return dummy.Next
}
