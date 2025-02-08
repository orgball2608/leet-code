package medium

func searchRange(nums []int, target int) []int {
	lower := searchLower(nums, target)
	upper := searchUpper(nums, target) - 1

	if lower <= upper {
		return []int{lower, upper}
	}
	return []int{-1, -1}
}

func searchLower(nums []int, target int) int {
	lo, hi := 0, len(nums)
	for lo < hi {
		mid := lo + (hi-lo)/2
		if nums[mid] < target {
			lo = mid + 1
		} else {
			hi = mid
		}
	}
	return lo
}

func searchUpper(nums []int, target int) int {
	lo, hi := 0, len(nums)
	for lo < hi {
		mid := lo + (hi-lo)/2
		if nums[mid] <= target {
			lo = mid + 1
		} else {
			hi = mid
		}
	}
	return lo
}

// Time: O(logN)
// Space: O(1)
