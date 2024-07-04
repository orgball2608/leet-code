package medium

func productExceptSelf(nums []int) []int {
	length := len(nums)
	result := make([]int, length)
	pre := 1
	for i, val := range nums {
		result[i] = pre
		pre = val * pre
	}
	pos := 1
	for i := length - 1; i >= 0; i-- {
		result[i] *= pos
		pos = nums[i] * pos
	}
	return result
}
