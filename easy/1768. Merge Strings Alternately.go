package easy

import "strings"

func mergeAlternately(word1 string, word2 string) string {
	count := 0
	var result strings.Builder
	for count < len(word1) || count < len(word2) {
		if count < len(word1) {
			result.WriteByte(word1[count])
		}
		if count < len(word2) {
			result.WriteByte(word2[count])
		}
		count++
	}
	return result.String()
}

// Time: O(N+M)
// Space: O(N+M)
