package medium

import "slices"

func successfulPairs(spells []int, potions []int, success int64) []int {
	slices.Sort(potions)
	succ := int(success) - 1
	pairs := make([]int, 0, len(spells))
	for i := range spells {
		atLeast := succ/spells[i] + 1
		pos, _ := slices.BinarySearch(potions, atLeast)
		pairs = append(pairs, len(potions)-pos)
	}
	return pairs
}

// Time: O(MlogM+NlogM)
// Space: O(N)
