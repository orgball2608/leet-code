package medium

func removeKdigits(num string, k int) string {
	if k >= len(num) {
		return "0"
	}

	minStack := []byte{}
	for i := range num {
		for k > 0 && len(minStack) > 0 && minStack[len(minStack)-1] > num[i] {
			minStack = minStack[:len(minStack)-1]
			k--
		}

		minStack = append(minStack, num[i])
	}

	minStack = minStack[:len(minStack)-k]

	result := strings.TrimLeft(string(minStack), "0")

	if result == "" {
		return "0"
	}

	return result
}

// Time: O(N)
// Space: O(N)
