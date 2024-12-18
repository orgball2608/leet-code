package medium

import "container/heap"

type Pair struct {
	Value int
	Freq  int
}

type MaxHeap []Pair

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i].Freq > h[j].Freq }
func (h MaxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(Pair))
}

func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func mostFrequentIDs(nums []int, freq []int) []int64 {
	h := &MaxHeap{}
	heap.Init(h)
	frequencyMap := map[int]int{}
	ans := make([]int64, len(nums))
	for i, val := range nums {
		frequencyMap[val] += freq[i]
		if frequencyMap[val] <= 0 {
			delete(frequencyMap, val)
		}

		heap.Push(h, Pair{
			Value: val,
			Freq:  frequencyMap[val],
		})

		for h.Len() > 0 {
			top := (*h)[0]
			if frequencyMap[top.Value] != top.Freq {
				heap.Pop(h)
			} else {
				break
			}
		}

		if h.Len() == 0 {
			ans[i] = 0
		} else {
			ans[i] = int64((*h)[0].Freq)
		}
	}

	return ans
}

// Time: O(NlogN)
// Space: O(N)
