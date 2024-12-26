package medium

import "sort"

func closeStrings(word1 string, word2 string) bool {
	if len(word1) != len(word2) {
		return false
	}

	set1 := make([]bool, 26)
	set2 := make([]bool, 26)
	counter1 := make([]int, 26)
	counter2 := make([]int, 26)
	for i := 0; i < len(word1); i++ {
		set2[word1[i]-'a'] = true
		set1[word2[i]-'a'] = true
		counter1[word1[i]-'a']++
		counter2[word2[i]-'a']++
	}

	for i := 0; i < 26; i++ {
		if set1[i] != set2[i] {
			return false
		}
	}

	sort.Ints(counter1)
	sort.Ints(counter2)

	for i := 0; i < 26; i++ {
		if counter1[i] != counter2[i] {
			return false
		}
	}

	return true
}

// Time: O(N)
// Space: O(1)
