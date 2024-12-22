package easy

func isSubsequence(s string, t string) bool {
	i, j := 0, 0
	for ; j < len(t) && i < len(s); j++ {
		if s[i] == t[j] {
			i++
		}
	}

	return i == len(s)
}
