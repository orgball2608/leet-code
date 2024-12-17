package medium

func eraseOverlapIntervals(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})

	curEnd := intervals[0][1]
	count := 0
	for i := 1; i < len(intervals); i++ {
		if curEnd > intervals[i][0] {
			count++
			continue
		}
		curEnd = intervals[i][1]

	}

	return count
}

// Time: O(NlogN)
// Space: O(1)
