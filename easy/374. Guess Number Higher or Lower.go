package easy

func guessNumber(n int) int {
	low, high := 1, n
	for {
		mid := low + (high-low)>>1
		check := guess(mid)
		if check == 0 {
			return mid
		}
		if check == 1 {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
}

// Time: O(logN)
// Space: O(1)
