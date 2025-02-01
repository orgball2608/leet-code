package medium

func checkSubarraySum(nums []int, k int) bool {
	remiderMap := map[int]int{
		0: -1,
	}

	prefixSum := 0
	for i, val := range nums {
		prefixSum += val
		remainder := prefixSum % k
		if remainder < 0 {
			remainder += k
		}

		prevIndex, ok := remiderMap[remainder]
		if !ok {
			remiderMap[remainder] = i
			continue
		}

		if i-prevIndex > 1 {
			return true
		}

	}

	return false
}

// Time: O(N)
// Space: O(N)
