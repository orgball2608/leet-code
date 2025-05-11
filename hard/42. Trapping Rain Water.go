package hard

func trap(height []int) int {
	n := len(height)
	stack := list.New()
	total := 0

	for i := n - 1; i >= 0; i-- {
		for stack.Len() > 0 && height[i] > height[stack.Back().Value.(int)] {
			bottom := stack.Back().Value.(int)
			stack.Remove(stack.Back())

			if stack.Len() == 0 {
				break
			}

			right := stack.Back().Value.(int)
			width := right - i - 1
			h := min(height[right], height[i]) - height[bottom]
			if h > 0 {
				total += width * h
			}
		}
		stack.PushBack(i)
	}

	return total
}

// Time: O(N)
// Space: O(N)
