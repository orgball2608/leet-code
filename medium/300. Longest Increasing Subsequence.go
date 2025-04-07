package medium

func lengthOfLIS(nums []int) int {
	L := make([]int, len(nums), len(nums))
	for i, _ := range nums {
		L[i] = 1
	}

	for i, val := range nums {
		for j := 0; j < i; j++ {
			if val > nums[j] {
				L[i] = max(L[j]+1, L[i])
			}
		}
	}

	max := 0
	for _, val := range L {
		if val > max {
			max = val
		}
	}

	return max
}

// Time: O(N*N)
// Space: O(N)

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

// Time: O(N*logN)
// Space: O(N)
