package medium

func singleNumber(nums []int) int {
	set := make(map[int]int)
	for _, val := range nums {
		set[val]++
	}

	for key, count := range set {
		if count == 1 {
			return key
		}
	}

	return -1
}

// Time: O(N)
// Space: O(N)
