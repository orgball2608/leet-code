package medium

func checkPossibility(nums []int) bool {
	count := 0
	for i := 1; i < len(nums); i++ {
		if nums[i-1] > nums[i] {
			count++
			if count > 1 {
				return false
			}

			if i == 1 || nums[i] >= nums[i-2] {
				nums[i-1] = nums[i]
			} else {
				nums[i] = nums[i-1]
			}
		}
	}

	return true
}

// Time: O(N)
// Space: O(1)
