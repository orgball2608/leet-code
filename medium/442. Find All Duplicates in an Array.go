package medium

func findDuplicates(nums []int) []int {
	set := make(map[int]struct{})
	var result []int
	for _, val := range nums {
		if _, ok := set[val]; ok {
			result = append(result, val)
			continue
		}

		set[val] = struct{}{}
	}

	return result
}

// Time: O(N)
// Space: O(N)
