package easy

func isAnagram(s string, t string) bool {
	hashmap := make(map[rune]int)

	for _, val := range s {
		hashmap[val]++
	}

	for _, val := range t {
		count, ok := hashmap[val]
		if ok {
			if count > 1 {
				hashmap[val] = count - 1
			} else {
				delete(hashmap, val)
			}
		} else {
			return false
		}
	}

	return len(hashmap) == 0
}

//Space: O(N)
//Time: O(N+M)
