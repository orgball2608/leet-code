package medium

func findRadius(houses []int, heaters []int) int {
	sort.Ints(heaters)
	ans := 0
	for _, val := range houses {
		index := sort.Search(len(heaters), func(i int) bool {
			return heaters[i] >= val
		})

		needRadius := math.MaxInt32
		if index < len(heaters) {
			needRadius = min(needRadius, heaters[index]-val)
		}

		if index-1 >= 0 {
			needRadius = min(needRadius, val-heaters[index-1])
		}

		ans = max(ans, needRadius)
	}

	return ans
}

// Time: O(NlogN + M*logN)
// Space: O(logN)
