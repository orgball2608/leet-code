package medium

func longestSubarray(nums []int) int {
	start, maxlength := 0, 0
	one := 0
	for i, value := range nums {
		if value == 1 {
			one++
		}
		if i-start+1-one > 1 {
			if nums[start] == 1 {
				one--
			}
			start++
		}
		if maxlength < (i - start + 1) {
			maxlength = i - start + 1
		}
	}
	// maxlength == find the longest substring with Ones after Replacement one zero
	// length = maxlength - 1
	return maxlength - 1

}
