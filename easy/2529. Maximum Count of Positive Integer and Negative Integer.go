package easy

func maximumCount(nums []int) int {
	left, right := 0, len(nums)-1
	zeroCount := 0
	for left <= right && nums[left] <= 0 {
		if nums[left] == 0 {
			zeroCount++
		}
		left++
	}

	negCount := left - zeroCount
	posCount := len(nums) - left
	return max(negCount, posCount)
}
