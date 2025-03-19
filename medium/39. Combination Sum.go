package medium

func combinationSum(candidates []int, target int) [][]int {
	var ans [][]int
	var currentSolutions []int
	var currentSum int

	var choose func(i int)
	choose = func(i int) {
		if i == len(candidates) {
			if currentSum == target {
				var solution []int
				for j, count := range currentSolutions {
					for k := 0; k < count; k++ {
						solution = append(solution, candidates[j])
					}
				}
				ans = append(ans, solution)
			}
			return
		}

		upperBound := (target - currentSum) / candidates[i]
		for j := 0; j <= upperBound; j++ {
			currentSolutions = append(currentSolutions, j)
			currentSum += j * candidates[i]
			choose(i + 1)
			currentSolutions = currentSolutions[:len(currentSolutions)-1]
			currentSum -= j * candidates[i]
		}
	}

	choose(0)
	return ans
}

// Time: O(K^N)
// Space: O(N) current_solution + O(N) recrussion stack
