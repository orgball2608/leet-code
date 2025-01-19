package medium

func thirdMax(nums []int) int {
	var first, second, third *int

	for _, num := range nums {
		if (first != nil && num == *first) ||
			(second != nil && num == *second) ||
			(third != nil && num == *third) {
			continue
		}

		if first == nil || num > *first {
			third = second
			second = first
			first = &num
		} else if second == nil || num > *second {
			third = second
			second = &num
		} else if third == nil || num > *third {
			third = &num
		}
	}

	if third != nil {
		return *third
	}

	return *first
}

// Time: O(N)
// Space: O(1)
