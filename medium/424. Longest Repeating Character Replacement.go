package medium

func characterReplacement(s string, k int) int {
	dict := make(map[byte]int, 26)
	start, maxlength, maxRepeatChar := 0, 0, 0
	for i, _ := range s {
		dict[s[i]]++
		if maxRepeatChar < dict[s[i]] {
			maxRepeatChar = dict[s[i]]
		}
		if (i - start + 1 - maxRepeatChar) > k {
			dict[s[start]]--
			start++
		}
		if maxlength < i-start+1 {
			maxlength = i - start + 1
		}
	}
	return maxlength
}
