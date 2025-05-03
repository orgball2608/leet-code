package easy

func lengthOfLastWord(s string) int {
	endWordIndex := len(s) - 1
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] != ' ' {
			endWordIndex = i
			break
		}
	}

	for i := endWordIndex; i >= 0; i-- {
		if s[i] == ' ' {
			return endWordIndex - i
		}
		if i == 0 {
			return endWordIndex - i + 1
		}
	}

	return 0
}

// Time: O(N)
// Space: O(1)
