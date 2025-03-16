package medium

func maxSumMinProduct(nums []int) int {
	stack := []int{}
	left := make([]int, len(nums))
	for i, val := range nums {
		for len(stack) > 0 && nums[stack[len(stack)-1]] >= val {
			stack = stack[:len(stack)-1]
		}

		if len(stack) > 0 {
			left[i] = stack[len(stack)-1] + 1
		} else {
			left[i] = 0
		}
		stack = append(stack, i)
	}

	stack = []int{}
	right := make([]int, len(nums))
	for i := len(nums) - 1; i >= 0; i-- {
		for len(stack) > 0 && nums[stack[len(stack)-1]] >= nums[i] {
			stack = stack[:len(stack)-1]
		}

		if len(stack) > 0 {
			right[i] = stack[len(stack)-1]
		} else {
			right[i] = len(nums)
		}
		stack = append(stack, i)
	}

	prefixSum := make([]int64, len(nums)+1)
	for i := 0; i < len(nums); i++ {
		prefixSum[i+1] = prefixSum[i] + int64(nums[i])
	}

	var maxProduct int64 = 0
	for i, val := range nums {
		totalSum := prefixSum[right[i]] - prefixSum[left[i]]
		product := totalSum * int64(val)
		if product > maxProduct {
			maxProduct = product
		}
	}

	const MOD = 1_000_000_007
	return int(maxProduct % MOD)
}

// Time: O(N)
// Space: O(N)
