package medium

func permuteUnique(nums []int) [][]int {
	sort.Ints(nums)
	curSolutions := make([]int, 0, len(nums))
	ans := [][]int{}
	used := make([]bool, len(nums))
	var backtrack func()
	backtrack = func() {
		if len(curSolutions) == len(nums) {
			temp := make([]int, len(nums))
			copy(temp, curSolutions)
			ans = append(ans, temp)
			return
		}

		for j := range nums {
			if used[j] {
				continue
			}

			if j > 0 && nums[j] == nums[j-1] && !used[j-1] {
				continue
			}

			used[j] = true
			curSolutions = append(curSolutions, nums[j])
			backtrack()
			curSolutions = curSolutions[:len(curSolutions)-1]
			used[j] = false
		}
	}

	backtrack()

	return ans
}

// Time: O(n*n!)
// Space: O(n)
