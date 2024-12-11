package medium

func lengthOfLIS(nums []int) int {
	L := make([]int, len(nums), len(nums))
	for i, _ := range nums {
		L[i] = 1
	}

	for i, val := range nums {
		for j := 0; j < i; j++ {
			if val > nums[j] {
				L[i] = max(L[j]+1, L[i])
			}
		}
	}

	max := 0
	for _, val := range L {
		if val > max {
			max = val
		}
	}

	return max
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

// Time: O(N*N)
// Space: O(N)
