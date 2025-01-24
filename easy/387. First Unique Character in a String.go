package easy

func firstUniqChar(s string) int {
	counter := make([]int, 26)
	for _, val := range s {
		counter[val-'a']++
	}

	for i, val := range s {
		count := counter[val-'a']
		if count == 1 {
			return i
		}
	}

	return -1
}

// Time: O(2N)
// Space: O(1)
