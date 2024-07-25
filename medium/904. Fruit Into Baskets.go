package medium

func totalFruit(fruits []int) int {
	n := len(fruits)
	if n < 3 {
		return n
	}
	backet := make(map[int]int)
	start, length := 0, 0
	for end := 0; end < n; end++ {
		backet[fruits[end]]++
		for len(backet) == 3 {
			backet[fruits[start]]--
			if backet[fruits[start]] == 0 {
				delete(backet, fruits[start])
			}
			start++
		}
		if length < (end - start + 1) {
			length = end - start + 1
		}
	}
	return length
}
