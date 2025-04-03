package medium

func subarraySum(nums []int, k int) int {
	counter := map[int]int{
		0: 1,
	}
	sum := 0
	count := 0
	for _, val := range nums {
		sum += val
		data, exists := counter[sum-k]
		if !exists {
			counter[sum]++
			continue
		}

		count += data
		counter[sum]++
	}
	return count
}

// Time: O(N)
// Space: O(N)
