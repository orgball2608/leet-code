package medium

func makesquare(matchsticks []int) bool {
	curSum := make([]int, 4)

	sum := 0
	for _, val := range matchsticks {
		sum += val
	}

	sort.Slice(matchsticks, func(i, j int) bool {
		return matchsticks[i] > matchsticks[j]
	})

	if sum%4 != 0 {
		return false
	}
	sideLen := sum / 4

	var choose func(i int) bool
	choose = func(i int) bool {
		if i == len(matchsticks) {
			for _, val := range curSum {
				if val != sideLen {
					return false
				}
			}
			return true
		}

		for j := 0; j < 4; j++ {
			if curSum[j]+matchsticks[i] == sideLen || curSum[j]+matchsticks[i]+matchsticks[len(matchsticks)-1] <= sideLen {
				curSum[j] += matchsticks[i]
				if choose(i + 1) {
					return true
				}
				curSum[j] -= matchsticks[i]
				if curSum[j] == 0 {
					break
				}
			}
		}
		return false
	}

	return choose(0)
}

// Time: O(4^N)
// Space: O(logN) sort + O(N) recursion
