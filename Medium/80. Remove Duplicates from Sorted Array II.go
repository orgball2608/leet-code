package medium

func removeDuplicates(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}
	next := 2
	for i := 2; i < len(nums); i++ {
		if nums[next-2] != nums[i] {
			nums[next] = nums[i]
			next++
		}
	}
	return next
}
