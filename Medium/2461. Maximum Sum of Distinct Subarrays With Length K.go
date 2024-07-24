package medium

func maximumSubarraySum(nums []int, k int) int64 {
	n := len(nums)
	if k <= 0 || k > n {
		return 0
	}
	frequency := make(map[int]int)
	start, sum, maxSum := 0, 0, 0
	for i, val := range nums {
		frequency[val]++
		sum += val
		if i-start+1 == k {
			if len(frequency) == k {
				maxSum = max(maxSum, sum)
			}

			if frequency[nums[start]] == 1 {
				delete(frequency, nums[start])
			} else {
				frequency[nums[start]]--
			}
			sum -= nums[start]
			start++
		}
	}

	return int64(maxSum)
}
