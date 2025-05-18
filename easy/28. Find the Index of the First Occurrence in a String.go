package easy

func strStr1(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}
	for i := 0; i <= len(haystack)-len(needle); i++ {
		match := true
		for j := 0; j < len(needle); j++ {
			if haystack[i+j] != needle[j] {
				match = false
			}
		}
		if match {
			return i
		}
	}
	return -1
}

// Time: O(N*M)
// Space: 0(1)

func strStr2(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}

	base := 256
	mod := 1_000_000_007
	m, n := len(needle), len(haystack)

	if m > n {
		return -1
	}

	hashNeedle := 0
	for i := 0; i < m; i++ {
		hashNeedle = (hashNeedle*base + int(needle[i])) % mod
	}

	power := 1
	for i := 0; i < m-1; i++ {
		power = (power * base) % mod
	}

	hashHay := 0
	for i := 0; i < m; i++ {
		hashHay = (hashHay*base + int(haystack[i])) % mod
	}

	if hashHay == hashNeedle && haystack[0:m] == needle {
		return 0
	}

	for i := m; i < n; i++ {
		hashHay = (hashHay - int(haystack[i-m])*power%mod + mod) % mod
		hashHay = (hashHay*base + int(haystack[i])) % mod

		if hashHay == hashNeedle && haystack[i-m+1:i+1] == needle {
			return i - m + 1
		}
	}

	return -1
}

// Time: O(N)
// Space: O(1)
