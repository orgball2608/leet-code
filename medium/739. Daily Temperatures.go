package medium

func dailyTemperatures(temperatures []int) []int {
	stack := Stack[int]{}
	ans := make([]int, len(temperatures))
	for i := 0; i < len(temperatures); i++ {
		for !stack.IsEmpty() && temperatures[i] > temperatures[stack.Peek()] {
			idx := stack.Pop()
			ans[idx] = i - idx
		}
		stack.Push(i)
	}

	return ans
}

// Time: O(n)
// Space: O(n)
