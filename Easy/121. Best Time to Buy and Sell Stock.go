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

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//Time: O(N)
//Space: O(1)
