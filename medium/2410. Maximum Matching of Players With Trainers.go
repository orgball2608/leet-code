package medium

func matchPlayersAndTrainers(players []int, trainers []int) int {
	sort.Ints(players)
	sort.Ints(trainers)
	i := 0

	for j := 0; i < len(players) && j < len(trainers); j++ {
		if trainers[j] >= players[i] {
			i++
		}
	}
	return i
}

// Time: O(NlogN)
// Space: O(1)
