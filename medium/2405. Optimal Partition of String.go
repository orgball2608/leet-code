package medium

func partitionString(s string) int {
	set := make([]bool, 26)
	res := 1
	for _, val := range s {
		ok := set[val-'a']
		if ok {
			res++
			set = make([]bool, 26)
		}
		set[val-'a'] = true
	}
	return res
}

// Time: O(N)
// Space: O(1)
