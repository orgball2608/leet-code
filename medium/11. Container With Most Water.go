package medium

func maxArea(height []int) int {
	left, right, maxArea := 0, len(height)-1, 0
	for left < right {
		minHeight := 0
		if height[left] < height[right] {
			minHeight = height[left]
		} else {
			minHeight = height[right]
		}

		if maxArea < (minHeight * (right - left)) {
			maxArea = minHeight * (right - left)
		}

		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return maxArea
}

//Time: O(N)
//Space: O(1)
