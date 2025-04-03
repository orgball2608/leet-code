package medium

func customSortString(order string, s string) string {
	counter := make([]int, 26)
	for i := 0; i < len(s); i++ {
		counter[s[i]-'a']++
	}

	var buffer bytes.Buffer
	for i := 0; i < len(order); i++ {
		count := counter[order[i]-'a']
		if count == 0 {
			continue
		}

		for j := 0; j < count; j++ {
			buffer.WriteByte(order[i])
		}
		counter[order[i]-'a'] = 0
	}

	for key, count := range counter {
		for i := 0; i < count; i++ {
			buffer.WriteByte(byte(key) + 'a')
		}
	}

	return buffer.String()
}

// Time: O(N + M)
// Space: O(1)
