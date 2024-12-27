package easy

type RecentCounter struct {
	Queue []int
}

func Constructor() RecentCounter {
	return RecentCounter{
		Queue: []int{},
	}
}

func (this *RecentCounter) Ping(t int) int {
	min := t - 3000
	this.Queue = append(this.Queue, t)
	for len(this.Queue) > 0 && this.Queue[0] < min {
		this.Queue = this.Queue[1:]
	}

	return len(this.Queue)
}

/**
 * Your RecentCounter object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Ping(t);
 */

// Time: O(1)
// Space: O(N)
