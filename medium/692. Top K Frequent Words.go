package medium

func topKFrequent(words []string, k int) []string {
	counter := make(map[string]int)
	for _, val := range words {
		counter[val]++
	}

	type Pair struct {
		Value string
		Count int
	}
	pairs := make([]Pair, 0, len(counter))
	for k, c := range counter {
		pairs = append(pairs, Pair{
			Value: k,
			Count: c,
		})
	}

	sort.Slice(pairs, func(i, j int) bool {
		if pairs[i].Count == pairs[j].Count {
			return pairs[i].Value < pairs[j].Value
		}

		return pairs[i].Count > pairs[j].Count
	})

	results := make([]string, 0, k)
	for i := 0; i < k && i < len(pairs); i++ {
		results = append(results, pairs[i].Value)
	}

	return results
}

// Time: O(N+M+MlogM+K)
// Space: O(M+M+K)
