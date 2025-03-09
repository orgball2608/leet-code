package medium

func rangeBitwiseAnd(left int, right int) int {
	result := left
	for i := left + 1; i <= right; i++ {
		result &= i
		if result == 0 {
			return 0
		}
	}
	return result
}

// Time: O(N)
// Space: O(1)

func rangeBitwiseAnd(left int, right int) int {
	for left < right {
		right &= (right - 1)
	}
	return right
}

// Time: O(logN)
// Space: O(1)
