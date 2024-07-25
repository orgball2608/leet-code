package medium

func numOfSubarrays(arr []int, k int, threshold int) int {
	l, sum, res := 0, 0, 0
	for r, val := range arr {
		sum += val
		if r-l+1 >= k {
			if sum/k >= threshold {
				res++
			}
			sum -= arr[l]
			l++
		}
	}
	return res
}
