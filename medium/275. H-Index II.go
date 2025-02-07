package medium

func hIndex(citations []int) int {
	lo, hi := 0, len(citations)
	for lo < hi {
		mid := lo + (hi-lo+1)/2
		idx := sort.Search(len(citations), func(i int) bool {
			return citations[i] >= mid
		})

		if len(citations)-idx >= mid {
			lo = mid
		} else {
			hi = mid - 1
		}
	}

	return lo
}

// Time: O((logN)^2)
// Space: O(1)
