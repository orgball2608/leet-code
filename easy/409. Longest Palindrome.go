package easy

func longestPalindrome(s string) int {
	count := [128]int{}

	for i := 0; i < len(s); i++ {
		count[s[i]]++
	}

	length := 0
	hasOdd := false
	for i := 0; i < 128; i++ {
		if count[i]%2 == 0 {
			length += count[i]
		} else {
			length += count[i] - 1
			hasOdd = true
		}
	}

	if hasOdd {
		length++
	}

	return length
}

// Time Complexity: O(n)
// Space Complexity: O(1)
