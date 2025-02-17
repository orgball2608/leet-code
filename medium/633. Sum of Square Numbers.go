package medium

func judgeSquareSum(c int) bool {
	left := 0
	right := int(math.Round(math.Sqrt(float64(c))))
	for left <= right {
		sum := left*left + right*right
		if sum == c {
			return true
		} else if sum < c {
			left++
		} else {
			right--
		}
	}

	return false
}

// Time: O(N)
// Space: O(1)
