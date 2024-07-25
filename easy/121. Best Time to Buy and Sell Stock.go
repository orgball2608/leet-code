package easy

func maxProfit(prices []int) int {
	length := len(prices)
	if length == 1 {
		return 0
	}
	left, right := 0, 1
	profit := 0
	for right < length {
		curProfit := prices[right] - prices[left]
		if curProfit < 0 {
			left = right
		} else {
			profit = max(profit, curProfit)
		}
		right++
	}
	return profit
}

//Time: O(N)
//Space: O(1)
