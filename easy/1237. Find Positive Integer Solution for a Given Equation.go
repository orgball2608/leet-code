package easy

/**
 * This is the declaration of customFunction API.
 * @param  x    int
 * @param  x    int
 * @return 	    Returns f(x, y) for any given positive integers x and y.
 *			    Note that f(x, y) is increasing with respect to both x and y.
 *              i.e. f(x, y) < f(x + 1, y), f(x, y) < f(x, y + 1)
 */

func findSolution(customFunction func(int, int) int, z int) [][]int {
	var results [][]int
	x, y := 1, z
	for x <= z && y >= 1 {
		res := customFunction(x, y)
		if res == z {
			results = append(results, []int{x, y})
			x++
			y--
		} else if res > z {
			y--
		} else {
			x++
		}
	}
	return results
}

// Time: O(Z*2)
// Space: O(N)
