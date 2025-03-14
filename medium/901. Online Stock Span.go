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
