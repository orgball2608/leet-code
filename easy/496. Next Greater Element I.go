package easy

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	dict := make(map[int]int)
	var stack []int
	for i := len(nums2) - 1; i >= 0; i-- {
		for len(stack) != 0 && stack[len(stack)-1] <= nums2[i] {
			stack = stack[:len(stack)-1]
		}

		if len(stack) == 0 {
			dict[nums2[i]] = -1
		} else {
			dict[nums2[i]] = stack[len(stack)-1]
		}

		stack = append(stack, nums2[i])
	}

	res := make([]int, len(nums1))
	for i, num := range nums1 {
		res[i] = dict[num]
	}

	return res
}

// Time: O(N)
// Space: O(N)
