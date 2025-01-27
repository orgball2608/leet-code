package medium

func sortColors(nums []int) {
	counter := make([]int, 3)
	for _, val := range nums {
		counter[val]++
	}

	index := 0
	for val, count := range counter {
		for count > 0 {
			nums[index] = val
			index++
			count--
		}
	}
}

// Time: O(N)
// Space: O(1)
