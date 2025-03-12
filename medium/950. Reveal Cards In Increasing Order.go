package medium

func deckRevealedIncreasing(deck []int) []int {
	n := len(deck)
	if n == 0 {
		return []int{}
	}

	sort.Ints(deck)
	queue := make([]int, 0, n)
	for i := 0; i < n; i++ {
		queue = append(queue, i)
	}

	result := make([]int, n)
	for _, card := range deck {
		pos := queue[0]
		queue = queue[1:]
		result[pos] = card
		if len(queue) > 0 {
			queue = append(queue, queue[0])
			queue = queue[1:]
		}
	}

	return result
}

// Time: O(NlogN)
// Space: O(N)
