package easy

func findWords(words []string) []string {
	s1 := "qwertyuiop"
	s2 := "asdfghjkl"
	s3 := "zxcvbnm"
	hashMap := make(map[rune]int, 26)
	for _, val := range s1 {
		hashMap[val] = 1
	}

	for _, val := range s2 {
		hashMap[val] = 2
	}

	for _, val := range s3 {
		hashMap[val] = 3
	}

	ans := []string{}
	for _, str := range words {
		isSameLine := true
		for i := 0; i < len(str)-1; i++ {
			first := unicode.ToLower(rune(str[i]))
			curr := unicode.ToLower(rune(str[i+1]))
			if hashMap[curr] != hashMap[first] {
				isSameLine = false
				break
			}
		}
		if isSameLine {
			ans = append(ans, str)
		}
	}

	return ans

}

// Time: O(N*M)
// Space: O(N)
