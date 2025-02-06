package medium

func isPerfectSquare(num int) bool {
	lo, hi := 1, num+1
	for lo < hi {
		mid := lo + (hi-lo)/2
		if mid*mid >= num {
			hi = mid
		} else {
			lo = mid + 1
		}
	}

	return lo*lo == num
}

// Time: O(LogN)
// Space: O(1)
