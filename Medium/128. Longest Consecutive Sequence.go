package medium

func longestConsecutive(nums []int) int {
	set := map[int]bool{}
	for _, val := range nums {
		set[val] = true
	}

	count := 0
	for _, val := range nums {
		if set[val+1] {
			continue
		}

		cnt := 0
		for curr := val; set[curr]; curr-- {
			cnt++
		}

		if cnt > count {
			count = cnt
		}
	}

	return count
}

//Time: O(N)
//Space: O(N)
