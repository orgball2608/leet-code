package easy

func getFinalState(nums []int, k int, multiplier int) []int {
	for i := 0; i < k; i++ {
		minIdx := 0
		for j := 1; j < len(nums); j++ {
			if nums[j] < nums[minIdx] {
				minIdx = j
			}
		}

		nums[minIdx] *= multiplier
	}

	return nums
}

// Time: O(K*N)
// Space: O(1)
