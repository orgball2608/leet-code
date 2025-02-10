package medium

func maximumUniqueSubarray(nums []int) int {
	left := 0
	set := make(map[int]int)
	maxSum, currentSum := 0, 0

	for right := 0; right < len(nums); right++ {
		if index, ok := set[nums[right]]; ok && index >= left {
			maxSum = max(maxSum, currentSum)

			for left <= index {
				currentSum -= nums[left]
				delete(set, nums[left])
				left++
			}
		}

		set[nums[right]] = right
		currentSum += nums[right]
	}

	return max(maxSum, currentSum)
}

// Time: O(N)
// Space: O(K)
