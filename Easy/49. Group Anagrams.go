package easy

func groupAnagrams(strs []string) [][]string {
	anagramMap := make(map[[26]int][]string)
	for _, val := range strs {
		var count [26]int
		for _, s := range val {
			count[s-'a']++
		}
		anagramMap[count] = append(anagramMap[count], val)
	}

	result := make([][]string, len(anagramMap))
	idx := 0

	for _, val := range anagramMap {
		result[idx] = val
		idx++
	}

	return result
}

//O(N*K)
//O(N*K)
