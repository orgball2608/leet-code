package medium

func findPeakElement(nums []int) int {
	start, end := 0, len(nums)-1
	for start < end {
		mid := start + (end-start)/2
		if nums[mid] > nums[mid+1] {
			end = mid
		} else {
			start = mid + 1
		}
	}

	return start
}

// Time: O(logN)
// Space: O(1)
