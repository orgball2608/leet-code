package easy

func findRelativeRanks(score []int) []string {
	hashMap := make(map[int]int, len(score))
	for i, val := range score {
		hashMap[val] = i
	}

	sort.Slice(score, func(i, j int) bool {
		return score[i] > score[j]
	})

	results := make([]string, len(score))
	for i := 0; i < len(score); i++ {
		index := hashMap[score[i]]
		var temp string
		switch i + 1 {
		case 1:
			temp = "Gold Medal"
		case 2:
			temp = "Silver Medal"
		case 3:
			temp = "Bronze Medal"
		default:
			temp = strconv.Itoa(i + 1)
		}

		results[index] = temp
	}

	return results
}

// Time: O(2N+NlogN)
// Space: O(2N)
