package easy

func isPowerOfTwo(n int) bool {
	if n < 1 {
		return false
	}

	for n%2 == 0 {
		n = n / 2
	}

	return n == 1
}

// Time: O(logN)
// Space: (1)

func isPowerOfTwo(n int) bool {
	if n < 1 {
		return false
	}

	if n == 1 {
		return true
	}

	if n%2 == 1 {
		return false
	}

	return isPowerOfTwo(n / 2)
}

// Time: O(logN)
// Space: (logN)

func isPowerOfTwo(n int) bool {
	if n == 0 {
		return false
	}
	return (n-1)&n == 0
}

// Time: O(1)
// Space: O(1)

func isPowerOfTwo(n int) bool {
	if n < 1 {
		return false
	}

	return (n & -n) == n
}

// Time: O(1)
// Space: O(1)
