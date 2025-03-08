package medium

func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)
	set := make(map[string]struct{})
	len := len(nums)
	subsets := make([][]int, 0, 1<<len)
	for i := 0; i < 1<<len; i++ {
		subset := make([]int, 0, len)
		var key string
		for j := 0; j < len; j++ {
			bit := (i >> j) & 1
			if bit == 1 {
				subset = append(subset, nums[j])
				key += strconv.Itoa(nums[j])
			}
		}

		if _, exists := set[key]; !exists {
			subsets = append(subsets, subset)
			set[key] = struct{}{}
		}
	}

	return subsets
}

// Time: O(N*2^N)
// Space: O(N*2^N)
