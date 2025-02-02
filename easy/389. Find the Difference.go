package easy

func findTheDifference(s string, t string) byte {
	sumS, sumT := 0, 0

	for _, ch := range s {
		sumS += int(ch)
	}

	for _, ch := range t {
		sumT += int(ch)
	}

	return byte(sumT - sumS)
}

// Time: O(N+M)
// Space: O(1)
