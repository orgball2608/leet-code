package medium

func triangleNumber(nums []int) int {
	sort.Ints(nums)
	ans := 0
	for i := len(nums) - 1; i >= 2; i-- {
		left, right := 0, i-1
		for left < right {
			if nums[left]+nums[right] > nums[i] {
				ans += right - left
				right--
			} else {
				left++
			}
		}
	}

	return ans
}

// Time: O(NlogN + N)
// Space: O(1) with out sort space
