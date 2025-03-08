package medium

func subsets(nums []int) [][]int {
	subsets := make([][]int, 0, 1<<len(nums))
	for i := 0; i < (1 << len(nums)); i++ {
		subset := make([]int, 0, len(nums))
		for j := 0; j < len(nums); j++ {
			bit := (i >> j) & 1
			if bit == 1 {
				subset = append(subset, nums[j])
			}
		}
		subsets = append(subsets, subset)
	}
	return subsets
}

// Time: O(N*2^N)
// Space: O(N * 2^N)
