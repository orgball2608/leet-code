package medium

func integerReplacement(n int) int {
	if n == 1 {
		return 0
	}

	if n%2 == 0 {
		return 1 + integerReplacement(n>>1)
	}

	return 1 + min(integerReplacement(n+1), integerReplacement(n-1))
}

// Time: O(logN)
// Space: O(logN)

func integerReplacement2(n int) int {
	steps := 0
	for n > 1 {
		if n&1 == 0 {
			n >>= 1
		} else if n == 3 || n%4 == 1 {
			n -= 1
		} else {
			n += 1
		}
		steps++
	}
	return steps
}

// Time: O(logN)
// Space: O(1)
