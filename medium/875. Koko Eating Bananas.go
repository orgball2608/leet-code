package medium

func minEatingSpeed(piles []int, h int) int {
	largestK := piles[0]
	for _, val := range piles {
		if val > largestK {
			largestK = val
		}
	}

	if h == len(piles) {
		return largestK
	}

	left, right := 1, largestK
	for left < right {
		mid := (left + right) / 2
		hours := h
		for i := 0; hours >= 0 && i < len(piles); i++ {
			hours -= (piles[i] + mid - 1) / mid
		}

		if hours < 0 {
			left = mid + 1
		} else {
			right = mid
		}
	}

	return right
}
