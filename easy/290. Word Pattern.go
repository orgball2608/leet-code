package easy

func wordPattern(pattern string, s string) bool {
	words := strings.Split(s, " ")
	if len(pattern) != len(words) {
		return false
	}

	pToW := make(map[byte]string)
	wToP := make(map[string]byte)
	for i, word := range words {
		w, ok := pToW[pattern[i]]
		if ok && w != word {
			return false
		}
		pToW[pattern[i]] = word

		p, ok := wToP[word]
		if ok && p != pattern[i] {
			return false
		}
		wToP[word] = pattern[i]
	}

	return true
}

// Time: O(N)
// Space: O(N)
