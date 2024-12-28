package easy

func gcdOfStrings(str1 string, str2 string) string {
	gcd := func(a, b int) int {
		for b != 0 {
			a, b = b, a%b
		}
		return a
	}

	if str1+str2 != str2+str1 {
		return ""
	}

	gcdLength := gcd(len(str1), len(str2))

	return str1[:gcdLength]
}

// Time: O(N+M)
// Space: O(1)
