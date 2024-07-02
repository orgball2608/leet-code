package easy

func backspaceCompare(s string, t string) bool {
	l1, l2 := len(s)-1, len(t)-1
	for l1 >= 0 || l2 >= 0 {
		i1 := getNextAvaibleIndex(s, l1)
		i2 := getNextAvaibleIndex(t, l2)
		if i1 < 0 && i2 < 0 {
			return true
		}
		if i1 < 0 || i2 < 0 || s[i1] != t[i2] {
			return false
		}
		l1 = i1 - 1
		l2 = i2 - 1
	}
	return l1 < 0 && l2 < 0
}

func getNextAvaibleIndex(s string, index int) int {
	countBackspace := 0
	for index >= 0 {
		if s[index] == '#' {
			countBackspace++
		} else if countBackspace > 0 {
			countBackspace--
		} else {
			break
		}
		index--
	}
	return index
}
