package medium

func myPow(x float64, n int) float64 {
	memo := make(map[int]float64)

	var f func(n int) float64
	f = func(n int) float64 {
		if n == 0 {
			return 1.0
		}

		if val, ok := memo[n]; ok {
			return val
		}

		if n < 0 {
			return 1.0 / f(-n)
		}

		half := f(n / 2)
		var result float64
		if n%2 == 1 {
			result = half * half * x
		} else {
			result = half * half
		}

		memo[n] = result
		return result
	}

	return f(n)
}

// Time: O(log n)
// Space: O(N for memoization) + O(log n for recursion stack) = O(N)
