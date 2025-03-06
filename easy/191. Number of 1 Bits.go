package easy

func hammingWeight(n int) int {
	count := 0
	for n > 0 {
		n = n & (n - 1)
		count++
	}

	return count
}

// Time: O(K) K is bit 1 of n
// Space: O(1)
