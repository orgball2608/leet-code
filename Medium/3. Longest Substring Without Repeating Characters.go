package medium

func lengthOfLongestSubstring(s string) int {
	dict := make(map[byte]int)
	start := 0
	length := 0
	for i, _ := range s {
		dict[s[i]]++
		for dict[s[i]] > 1 {
			if dict[s[start]] == 1 {
				delete(dict, s[start])
			} else {
				dict[s[start]]--
			}
			start++
		}
		if length < (i - start + 1) {
			length = (i - start + 1)
		}
	}
	return length
}
