package easy

func largestAltitude(gain []int) int {
	if len(gain) == 0 {
		return 0
	}

	hightAltitude := 0
	currentAltitude := 0
	for i := 0; i < len(gain); i++ {
		currentAltitude = currentAltitude + gain[i]
		hightAltitude = max(hightAltitude, currentAltitude)
	}

	return hightAltitude
}

// Time: O(N)
// Space: O(1)
