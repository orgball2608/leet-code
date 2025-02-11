package easy

func singleNumber(nums []int) int {
	var result int
	for _, val := range nums {
		result ^= val
	}
	return result
}

// Time: O(N)
// Space: O(1)
