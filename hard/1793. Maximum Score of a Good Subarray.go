package hard

func maximumScore(nums []int, k int) int {
	n := len(nums)
	stack := make([]int, 0, n)
	left := make([]int, n)
	for i := range nums {
		for len(stack) > 0 && nums[stack[len(stack)-1]] >= nums[i] {
			stack = stack[:len(stack)-1]
		}

		if len(stack) > 0 {
			left[i] = stack[len(stack)-1]
		} else {
			left[i] = -1
		}
		stack = append(stack, i)
	}

	stack = []int{}
	right := make([]int, n)
	for i := n - 1; i >= 0; i-- {
		for len(stack) > 0 && nums[stack[len(stack)-1]] >= nums[i] {
			stack = stack[:len(stack)-1]
		}

		if len(stack) > 0 {
			right[i] = stack[len(stack)-1]
		} else {
			right[i] = n
		}
		stack = append(stack, i)
	}

	result := 0
	for i := range nums {
		if k > left[i] && k < right[i] {
			result = max(result, (right[i]-left[i]-1)*nums[i])
		}
	}

	return result
}

// Time: 0(N)
// Space: 0(N)
