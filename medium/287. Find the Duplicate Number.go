package medium

func findDuplicate(nums []int) int {
	for _, val := range nums {
		index := abs(val) - 1
		if nums[index] < 0 {
			return abs(val)
		}
		nums[index] = -nums[index]
	}
	return 0
}

//Time: O(N)
//Space: O(1)
