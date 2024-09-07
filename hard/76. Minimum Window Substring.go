package hard

import "math"

func minWindow(s string, t string) string {
	if len(s) == 0 || len(t) == 0 || len(s) < len(t) {
		return ""
	}

	mapStr := make([]int, 128)
	count := len(t)
	start, end := 0, 0
	minLen, startIndex := math.MaxInt32, 0
	for _, val := range t {
		mapStr[val]++
	}

	for end < len(s) {
		if mapStr[s[end]] > 0 {
			count--
		}
		mapStr[s[end]]--
		end++

		for count == 0 {
			if end-start < minLen {
				minLen = end - start + 1
				startIndex = start
			}

			if mapStr[s[start]] == 0 {
				count++
			}
			mapStr[s[start]]++
			start++
		}
	}

	if minLen == math.MaxInt32 {
		return ""
	}

	return s[startIndex : startIndex+minLen-1]
}

//Time: 0(N)
//Space: O(1)
