package medium

func longestOnes(nums []int, k int) int {
	left, maxLen, zero := 0, 0, 0
	for right := 0; right < len(nums); right++ {
		if nums[right] == 0 {
			zero++
		}

		for zero > k {
			if nums[left] == 0 {
				zero--
			}
			left++
		}

		maxLen = max(maxLen, right-left+1)
	}

	return maxLen
}

// Time: O(N)
// Space: O(1)
