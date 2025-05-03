package easy

func majorityElement(nums []int) int {
	counter := make(map[int]int)
	for _, val := range nums {
		count, exists := counter[val]
		if exists {
			counter[val]++
		} else {
			counter[val] = 1
		}

		if count >= len(nums)/2 {
			return val
		}
	}

	return 0
}

// Time: O(N)
// Space: O(N)
