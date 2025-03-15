package medium

type StockSpanner struct {
	Data []int
}

func Constructor() StockSpanner {
	return StockSpanner{
		Data: []int{},
	}
}

func (this *StockSpanner) Next(price int) int {
	this.Data = append(this.Data, price)
	count := 0
	for len(this.Data)-count > 0 && this.Data[len(this.Data)-count-1] <= price {
		count++
	}

	return count
}

// Time: O(N)
// Space: O(N)

/**
 * Your StockSpanner object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Next(price);
 */

type StockSpanner struct {
	Queue [][2]int
}

func Constructor() StockSpanner {
	return StockSpanner{
		Queue: make([][2]int, 0),
	}
}

func (this *StockSpanner) Next(price int) int {
	span := 1
	for len(this.Queue) > 0 && this.Queue[len(this.Queue)-1][0] <= price {
		span += this.Queue[len(this.Queue)-1][1]
		this.Queue = this.Queue[:len(this.Queue)-1]
	}

	this.Queue = append(this.Queue, [2]int{price, span})

	return span
}

/**
 * Your StockSpanner object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Next(price);
 */

// Time: O(1)
// Space: O(N)
