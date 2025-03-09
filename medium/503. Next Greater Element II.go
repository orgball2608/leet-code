package medium

func nextGreaterElements(nums []int) []int {
	result := make([]int, len(nums))
	var stack []int
	n := len(nums)
	for i := 2*n - 1; i >= 0; i-- {
		for len(stack) > 0 && stack[len(stack)-1] <= nums[i%n] {
			stack = stack[:len(stack)-1]
		}

		if len(stack) > 0 {
			result[i%n] = stack[len(stack)-1]
		} else {
			result[i%n] = -1
		}

		stack = append(stack, nums[i%n])
	}

	return result
}

// Time: O(N)
// Space: O(N)
