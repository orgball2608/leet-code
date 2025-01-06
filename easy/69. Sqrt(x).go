package easy

func mySqrt(x int) int {
	left := 0
	right := 46341
	ans := 0
	for left <= right {
		mid := (right + left) / 2
		if mid*mid <= x {
			ans = mid
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return ans
}

// Time: O(logN)
// Space: O(1)
