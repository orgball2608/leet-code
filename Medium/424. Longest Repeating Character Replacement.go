package medium

func characterReplacement(s string, k int) int {
	dict := make(map[byte]int)
	start, maxlength, maxrepeatchar := 0, 0, 0
	for i, _ := range s {
		dict[s[i]]++
		if maxrepeatchar < dict[s[i]] {
			maxrepeatchar = dict[s[i]]
		}
		if (i - start + 1 - maxrepeatchar) > k {
			dict[s[start]]--
			start++
		}
		if maxlength < i-start+1 {
			maxlength = i - start + 1
		}
	}
	return maxlength
}
