package easy

func climbStairs(n int) int {
	var backtrack func(int) int
	dict := make(map[int]int)
	backtrack = func(num int) int {
		if v, ok := dict[num]; ok {
			return v
		}

		if num == 0 {
			return 1
		}

		if num < 0 {
			return 0
		}

		cur := backtrack(num-2) + backtrack(num-1)
		dict[num] = cur
		return cur
	}

	return backtrack(n)
}

// Time: O(N)
// Space: O(N)
