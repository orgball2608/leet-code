package medium

import "sort"

func threeSumMulti(arr []int, target int) int {
	const mod = 1000000007
	n := len(arr)
	count := 0

	sort.Ints(arr)

	for i := 0; i < n-2; i++ {
		left, right := i+1, n-1
		for left < right {
			currentSum := arr[i] + arr[left] + arr[right]
			if currentSum == target {
				if arr[left] == arr[right] {
					// If left and right are the same, we have a range of equal values.
					rangeCount := right - left + 1
					count = (count + (rangeCount * (rangeCount - 1) / 2)) % mod
					break
				} else {
					leftCount, rightCount := 1, 1

					// Count the number of elements equal to arr[left]
					for left+1 < right && arr[left] == arr[left+1] {
						left++
						leftCount++
					}

					// Count the number of elements equal to arr[right]
					for left < right-1 && arr[right] == arr[right-1] {
						right--
						rightCount++
					}

					count = (count + leftCount*rightCount) % mod
					left++
					right--
				}
			} else if currentSum < target {
				left++
			} else {
				right--
			}
		}
	}

	return count
}
