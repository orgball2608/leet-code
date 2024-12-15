package medium

func insert(intervals [][]int, newInterval []int) [][]int {
	preIntervals := [][]int{}
	postIntervals := [][]int{}

	for _, interval := range intervals {
		if interval[1] < newInterval[0] {
			preIntervals = append(preIntervals, interval)
		} else if interval[0] > newInterval[1] {
			postIntervals = append(postIntervals, interval)
		} else {
			newInterval[0] = min(newInterval[0], interval[0])
			newInterval[1] = max(newInterval[1], interval[1])
		}
	}

	result := append(preIntervals, newInterval)
	result = append(result, postIntervals...)
	return result
}

// Time: O(N)
// Space: O(N)
