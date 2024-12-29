package easy

func numberGame(nums []int) []int {
	slices.Sort(nums)
	for i := 1; i < len(nums); i += 2 {
		nums[i-1], nums[i] = nums[i], nums[i-1]
	}
	return nums
}

// Time: O(NlogN)
// Space: O(1)

//import (
//	"container/heap"
//)
//
//type MinHeap []int
//
//func (h MinHeap) Len() int           { return len(h) }
//func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
//func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
//
//func (h *MinHeap) Push(x any) {
//	*h = append(*h, x.(int))
//}
//
//func (h *MinHeap) Pop() any {
//	old := *h
//	n := len(old)
//	x := old[n-1]
//	*h = old[0 : n-1]
//	return x
//}
//
//func numberGame(nums []int) []int {
//	h := MinHeap(nums)
//	heap.Init(&h)
//
//	arr := []int{}
//	for h.Len() > 0 {
//		alicePick := heap.Pop(&h).(int)
//		bobPick := heap.Pop(&h).(int)
//
//		arr = append(arr, bobPick)
//		arr = append(arr, alicePick)
//	}
//
//	return arr
//}

// Time: O(NLogN)
// Space: O(N)
