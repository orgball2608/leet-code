package easy

func isPalindrome(s string) bool {
	left, right := 0, len(s)-1

	if s == " " {
		return true
	}

	for left <= right {
		if isPunctuation(s[left]) {
			left++
			continue
		}

		if isPunctuation(s[right]) {
			right--
			continue
		}

		if toLower(s[left]) != toLower(s[right]) {
			return false
		}

		left++
		right--
	}

	return true
}

func toLower(ch byte) byte {
	if ch >= 'A' && ch <= 'Z' {
		return ch + ('a' - 'A')
	}
	return ch
}

// check ascii table to validate
func isPunctuation(c byte) bool {
	if (c >= 32 && c <= 47) || (c >= 58 && c <= 64) || (c >= 91 && c <= 96) || (c >= 123 && c <= 126) {
		return true
	}
	return false
}

//Time: O(N)
//Space: O(1)
