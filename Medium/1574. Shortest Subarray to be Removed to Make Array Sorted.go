func findLengthOfShortestSubarray(arr []int) int {
	left, right := 0, len(arr)-1
	for left < len(arr)-1 && arr[left] <= arr[left+1] {
		left++
	}
	if left == len(arr)-1 {
		return 0
	}

	for right > 0 && arr[right] >= arr[right-1] {
		right--
	}
	res := min(right, len(arr)-left-1)

	i := 0
	j := right
	for i <= left && j <= len(arr)-1 {
		if arr[j] >= arr[i] {
			res = min(res, j-i-1)
			i++
		} else {
			j++
		}
	}

	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}