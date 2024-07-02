package easy

func reverseString(s []byte) {
	left, right := 0, len(s)-1
	for left < right {
		temp := s[right]
		s[right] = s[left]
		s[left] = temp
		right--
		left++
	}
}
