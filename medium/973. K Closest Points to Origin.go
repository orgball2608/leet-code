package medium

import "container/heap"

type Point struct {
	X int
	Y int
}

func (p *Point) DistanceToOrigin() float64 {
	return math.Sqrt(float64(p.X*p.X + p.Y*p.Y))
}

type MinHeap []Point

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i].DistanceToOrigin() < h[j].DistanceToOrigin() }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x any) {
	*h = append(*h, x.(Point))
}

func (h *MinHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func kClosest(points [][]int, k int) [][]int {
	h := &MinHeap{}
	heap.Init(h)
	for _, val := range points {
		heap.Push(h, Point{
			X: val[0],
			Y: val[1],
		})
	}

	result := [][]int{}
	for h.Len() > 0 && len(result) < k {
		temp := heap.Pop(h).(Point)
		result = append(result, []int{temp.X, temp.Y})
	}

	return result
}
