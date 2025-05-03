package medium

func reverseWords(s string) string {
	words := []string{}
	n := len(s)
	i := 0

	for i < n {
		for i < n && s[i] == ' ' {
			i++
		}
		if i >= n {
			break
		}

		start := i
		for i < n && s[i] != ' ' {
			i++
		}
		word := s[start:i]
		words = append(words, word)
	}

	left, right := 0, len(words)-1
	for left < right {
		words[left], words[right] = words[right], words[left]
		left = left + 1
		right = right - 1
	}

	return strings.Join(words, " ")
}

// Time: O(N)
// Space: O(N)
