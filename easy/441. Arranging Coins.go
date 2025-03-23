package easy

func arrangeCoins(n int) int {
	left := 1
	right := int(math.Sqrt(float64(2*n))) + 1
	ans := 0
	for left <= right {
		mid := left + (right-left)/2
		if calcCoinsNeeded(mid) <= n {
			ans = mid
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return ans
}

func calcCoinsNeeded(n int) int {
	return n * (n + 1) / 2
}

// Time: O(log(sqrt(2n)))
// Space: O(1)
