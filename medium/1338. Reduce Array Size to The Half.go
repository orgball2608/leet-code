package medium

import "slices"

func minSetSize(arr []int) int {
	counter := make([]int, 100001)
	for _, val := range arr {
		counter[val]++
	}

	slices.Sort(counter)
	slices.Reverse(counter)

	removed, ans := 0, 0
	half := len(arr) / 2
	for _, val := range counter {
		removed += val
		ans++
		if removed >= half {
			break
		}
	}

	return ans
}

// Time : O(N+100001log100001+N)
// Space: O(100001)
