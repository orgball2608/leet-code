package easy

import "slices"

func answerQueries(nums []int, queries []int) []int {
	results := make([]int, len(queries))
	slices.Sort(nums)
	for i := 1; i < len(nums); i++ {
		nums[i] += nums[i-1]
	}
	for i, val := range queries {
		idx, _ := slices.BinarySearch(nums, val+1)
		results[i] = idx
	}
	return results
}

//Time: O(NlogN+N+MlogN)
//Space: O(M)
