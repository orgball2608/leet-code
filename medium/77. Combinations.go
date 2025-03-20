package medium

func combine(n int, k int) [][]int {
	var currentSolution []int
	var result [][]int

	var backtrack func(i int)
	backtrack = func(i int) {
		if len(currentSolution) == k {
			result = append(result, append([]int{}, currentSolution...))
			return
		}

		remainingNeeded := k - len(currentSolution)
		for j := i; j <= n-remainingNeeded+1; j++ {
			currentSolution = append(currentSolution, j)
			backtrack(j + 1)
			currentSolution = currentSolution[:len(currentSolution)-1]
		}
	}

	backtrack(1)
	return result
}

// Time: O(K × N!/(K!(N-K)!))
// Space: O(K) recrussion stack
