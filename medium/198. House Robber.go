package medium

func rob(nums []int) int {
	prev1, prev2 := 0, 0
	for _, val := range nums {
		prev2, prev1 = prev1, max(prev1, prev2+val)
	}

	return prev1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Time: O(N)
// Space: O(1)
