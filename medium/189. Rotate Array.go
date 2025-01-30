package medium

func rotate(nums []int, k int) {
	requiredRotaion := k % len(nums)
	if requiredRotaion == 0 {
		return
	}

	reverse(nums)
	reverse(nums[:requiredRotaion])
	reverse(nums[requiredRotaion:])
}

func reverse(slice []int) []int {
	left := 0
	right := len(slice) - 1
	for left < right {
		slice[left], slice[right] = slice[right], slice[left]
		left++
		right--
	}

	return slice
}

// Time: O(N)
// Space: O(1)
