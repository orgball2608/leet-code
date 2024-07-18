package medium

func checkInclusion(s1 string, s2 string) bool {
	counter := [26]int{}
	left := 0
	for _, val := range s1 {
		counter[val-'a']++
	}

	for right, _ := range s2 {
		counter[s2[right]-'a']--
		if counter == [26]int{} {
			return true
		}

		if right-left+1 >= len(s1) {
			counter[s2[left]-'a']++
			left++
		}
	}

	return false
}

//Time: O(N+M)
//Space: O(1)
