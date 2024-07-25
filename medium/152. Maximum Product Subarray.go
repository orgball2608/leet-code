package medium

func maxProduct(nums []int) int {
	ans, curMax, curMin := nums[0], nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] < 0 {
			curMax, curMin = curMin, curMax
		}
		curMin = min(nums[i], nums[i]*curMin)
		curMax = max(nums[i], nums[i]*curMax)
		ans = max(curMax, ans)
	}
	return ans
}

//Time: O(N)
//Space: O(1)
