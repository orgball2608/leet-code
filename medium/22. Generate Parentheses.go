package medium

func generateParenthesis(n int) []string {
	res := make([]string, 0)
	backtrack(&res, "", 0, 0, n)
	return res
}

func backtrack(res *[]string, current string, left, right int, n int) {
	if len(current) == n*2 {
		*res = append(*res, current)
		return
	}

	if left < n {
		backtrack(res, current+"(", left+1, right, n)
	}

	if right < left {
		backtrack(res, current+")", left, right+1, n)
	}
}

// Time: O(4^n / sqrt(n))
// Space: O(4^n / sqrt(n))
