package medium

func findClosestElements(arr []int, k int, x int) []int {
	left, right := 0, len(arr)-1

	for right-left+1 > k {
		if abs(arr[left]-x) > abs(arr[right]-x) {
			left++
		} else {
			right--
		}
	}

	return arr[left : right+1]
}

func abs(a int) int {
	if a < 0 {
		return -a
	}

	return a
}
