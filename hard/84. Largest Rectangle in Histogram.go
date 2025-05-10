package hard

func largestRectangleArea(heights []int) int {
	n := len(heights)
	left := make([]int, n)
	right := make([]int, n)
	dq := list.New()

	for i := 0; i < n; i++ {
		for dq.Len() > 0 && heights[dq.Back().Value.(int)] >= heights[i] {
			dq.Remove(dq.Back())
		}

		if dq.Len() > 0 {
			left[i] = dq.Back().Value.(int)
		} else {
			left[i] = -1
		}
		dq.PushBack(i)
	}

	dq = list.New()
	for i := len(heights) - 1; i >= 0; i-- {
		for dq.Len() > 0 && heights[dq.Back().Value.(int)] >= heights[i] {
			dq.Remove(dq.Back())
		}

		if dq.Len() > 0 {
			right[i] = dq.Back().Value.(int)
		} else {
			right[i] = n
		}
		dq.PushBack(i)
	}

	maxArea := 0
	for i := range heights {
		width := right[i] - 1 - left[i] + 1 - 1
		area := width * heights[i]
		maxArea = max(maxArea, area)
	}

	return maxArea
}

// Time:O(N)
// Space: O(N)
