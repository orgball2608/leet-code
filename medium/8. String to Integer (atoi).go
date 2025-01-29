package medium

func myAtoi(s string) int {
	s = strings.TrimLeft(s, " ")
	if len(s) == 0 {
		return 0
	}

	sign := 1
	if s[0] == '-' {
		sign = -1
		s = s[1:]
	} else if s[0] == '+' {
		s = s[1:]
	}

	if len(s) == 0 {
		return 0
	}

	ans := 0
	for _, val := range s {
		if !unicode.IsDigit(val) {
			break
		}

		ans = ans*10 + int(val-'0')

		if ans*sign > 1<<31-1 {
			return 1<<31 - 1
		}

		if ans*sign < -1<<31 {
			return -1 << 31
		}
	}

	return ans * sign
}

// Time: O(N)
// Space: O(1)
