package medium

func shipWithinDays(weights []int, days int) int {
	low, hight := 0, 0
	for _, weight := range weights {
		low = max(low, weight)
		hight += weight
	}

	for low < hight {
		mid := low + (hight-low)/2
		if check(mid, weights, days) {
			hight = mid
		} else {
			low = mid + 1
		}
	}

	return low
}

func check(capacity int, weights []int, days int) bool {
	curcap := 0
	count := 1
	for _, weight := range weights {
		if curcap+weight <= capacity {
			curcap += weight
		} else {
			curcap = weight
			count++
		}
	}

	return count <= days
}

// Time: O(logN*N)
// Space: O(1)
