package medium

func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)

	var result [][]int
	var currentSolution []int

	var backtrack func(i int, sum int)
	backtrack = func(i int, sum int) {
		if sum == target {
			result = append(result, append([]int{}, currentSolution...))
			return
		}

		for j := i; j < len(candidates); j++ {
			if j > i && candidates[j] == candidates[j-1] {
				continue
			}

			if sum+candidates[j] > target {
				break
			}

			currentSolution = append(currentSolution, candidates[j])
			sum += candidates[j]
			backtrack(j+1, sum)
			currentSolution = currentSolution[:len(currentSolution)-1]
			sum -= candidates[j]
		}
	}

	backtrack(0, 0)
	return result
}

// Time: O(2^N)
// Space: O(logN) sort + O(N) recrussion stack
