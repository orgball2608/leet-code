package medium

import (
	"sort"
)

func frequencySort(s string) string {
	counter := make([]int, 128)
	for _, char := range s {
		counter[char]++
	}

	chars := []byte(s)
	sort.Slice(chars, func(i, j int) bool {
		if counter[chars[i]] == counter[chars[j]] {
			return chars[i] < chars[j]
		}
		return counter[chars[i]] > counter[chars[j]]
	})

	return string(chars)
}

// Time: O(N + NlogN + O(N))
// Space: O(N)

//import (
//	"sort"
//	"strings"
//)
//
//func frequencySort(s string) string {
//	counter := make(map[rune]int, 26)
//	for _, char := range s {
//		counter[char]++
//	}
//
//	type freqPair struct {
//		char rune
//		freq int
//	}
//	freqList := make([]freqPair, 0, len(counter))
//	for char, freq := range counter {
//		freqList = append(freqList, freqPair{char, freq})
//	}
//
//	sort.Slice(freqList, func(i, j int) bool {
//		return freqList[i].freq > freqList[j].freq
//	})
//
//	var sb strings.Builder
//	for _, freq := range freqList {
//		for freq.freq > 0 {
//			sb.WriteRune(rune(freq.char + 'a'))
//			c--
//		}
//	}
//
//	return sb.String()
//}

// Time: O(N + KlogK + O(N))
// Space: O(N)
