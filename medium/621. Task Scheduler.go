package medium

func leastInterval(tasks []byte, n int) int {
	counter := make(map[byte]int, 26)
	for _, val := range tasks {
		counter[val-'A']++
	}

	var maxCounter int
	for _, val := range counter {
		if val > maxCounter {
			maxCounter = val
		}
	}

	var maxSameCount int
	for _, val := range counter {
		if val == maxCounter {
			maxSameCount++
		}
	}

	res := (maxCounter-1)*(n+1) + maxSameCount
	return max(res, len(tasks))
}

// Time: O(N)
// Space: O(1)

//import "container/heap"
//
//type Counter struct {
//	Letter byte
//	Count  int
//}
//
//type MaxHeap []Counter
//
//func (h MaxHeap) Len() int           { return len(h) }
//func (h MaxHeap) Less(i, j int) bool { return h[i].Count > h[j].Count }
//func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
//
//func (h *MaxHeap) Push(x any) {
//	*h = append(*h, x.(Counter))
//}
//
//func (h *MaxHeap) Pop() any {
//	old := *h
//	n := len(old)
//	x := old[n-1]
//	*h = old[0 : n-1]
//	return x
//}
//
//func leastInterval(tasks []byte, n int) int {
//	counter := make(map[byte]int, 26)
//	for _, val := range tasks {
//		counter[val-'A']++
//	}
//
//	h := &MaxHeap{}
//	heap.Init(h)
//	for key, val := range counter {
//		heap.Push(h, Counter{
//			Letter: key,
//			Count:  val,
//		})
//	}
//
//	result := 0
//	for h.Len() > 0 {
//		rePush := make([]Counter, 0, h.Len())
//		time := 0
//		for i := 0; i <= n; i++ {
//			if h.Len() > 0 {
//				top := heap.Pop(h).(Counter)
//				time++
//				if top.Count-1 > 0 {
//					rePush = append(rePush, Counter{
//						Letter: top.Letter,
//						Count:  top.Count - 1,
//					})
//				}
//			}
//
//		}
//
//		if len(rePush) > 0 {
//			result += n + 1
//		} else {
//			result = result + time
//		}
//
//		for _, val := range rePush {
//			heap.Push(h, val)
//		}
//	}
//
//	return result
//}

// Time: O(NlogN)
// Space: O(N)
