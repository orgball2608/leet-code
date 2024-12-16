package medium

import "sort"

func merge(intervals [][]int) [][]int {
	if len(intervals) <= 1 {
		return intervals
	}

	sort.Slice(intervals, func(i int, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	left, right := intervals[0][0], intervals[0][1]
	result := [][]int{}
	for i := 1; i < len(intervals); i++ {
		if left <= intervals[i][1] && right >= intervals[i][0] {
			left = min(left, intervals[i][0])
			right = max(right, intervals[i][1])
			continue
		}
		result = append(result, []int{left, right})
		left = intervals[i][0]
		right = intervals[i][1]
	}
	result = append(result, []int{left, right})

	return result
}

// Time: O(NlogN)
// Space: O(N)
