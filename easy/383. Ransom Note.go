package easy

func canConstruct(ransomNote string, magazine string) bool {
	counter := make([]int, 26)

	for _, val := range magazine {
		counter[val-'a']++
	}

	for _, val := range ransomNote {
		counter[val-'a']--
		if counter[val-'a'] < 0 {
			return false
		}
	}

	return true
}

// Time: O(N+M) where N is the length of ransomNote and M is the length of magazine
// Space: O(1)
