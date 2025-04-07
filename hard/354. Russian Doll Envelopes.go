package hard

func maxEnvelopes(envelopes [][]int) int {
	sort.Slice(envelopes, func(i, j int) bool {
		if envelopes[i][0] < envelopes[j][0] {
			return true
		}

		if envelopes[i][0] == envelopes[j][0] {
			return envelopes[i][1] > envelopes[j][1]
		}

		return false
	})

	nums := make([]int, 0, len(envelopes))
	for _, val := range envelopes {
		nums = append(nums, val[1])
	}

	return lengthOfLIS(nums)
}

func lengthOfLIS(nums []int) int {
	var sub []int
	for _, val := range nums {
		if len(sub) == 0 || val > sub[len(sub)-1] {
			sub = append(sub, val)
			continue
		}

		i := sort.Search(len(sub), func(i int) bool { return sub[i] >= val })
		sub[i] = val
	}

	return len(sub)
}

// Time: O(NlogN)
// Space: O(logN) sort space complexity
