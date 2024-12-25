package easy

func pivotIndex(nums []int) int {
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}

	leftSum := 0
	for i := 0; i < len(nums); i++ {
		if leftSum == sum-nums[i]-leftSum {
			return i
		}
		leftSum += nums[i]
	}

	return -1
}

// Time: O(N)
// Space: O(1)
