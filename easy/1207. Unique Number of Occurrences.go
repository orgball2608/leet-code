package easy

func uniqueOccurrences(arr []int) bool {
	counter := make(map[int]int, len(arr))
	for _, val := range arr {
		counter[val]++
	}

	set := make(map[int]struct{}, len(counter))
	for _, val := range counter {
		if _, ok := set[val]; ok {
			return false
		}
		set[val] = struct{}{}
	}

	return true
}

// Time: O(N+P)
// Space: O(P+Q)
