package medium

type Car struct {
	pos  int
	time float64
}

func carFleet(target int, position []int, speed []int) int {
	n := len(position)
	cars := make([]Car, n)

	for i := 0; i < n; i++ {
		time := float64(target-position[i]) / float64(speed[i])
		cars[i] = Car{pos: position[i], time: time}
	}

	sort.Slice(cars, func(i, j int) bool {
		return cars[i].pos > cars[j].pos
	})

	fleets := 0
	lastTime := 0.0

	for _, car := range cars {
		if car.time > lastTime {
			fleets++
			lastTime = car.time
		}
	}

	return fleets
}

// Time: O(NlogN)
// Space: O(N) + O(Sort)
