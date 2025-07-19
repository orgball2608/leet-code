package medium

func numDecodings(s string) int {
	n := len(s)
	memo := make(map[int]int)

	var dfs func(i int) int
	dfs = func(i int) int {
		if i == n {
			return 1
		}
		if s[i] == '0' {
			return 0
		}
		if val, ok := memo[i]; ok {
			return val
		}

		res := dfs(i + 1)
		if i+1 < n {
			twoDigit := (s[i]-'0')*10 + (s[i+1] - '0')
			if twoDigit >= 10 && twoDigit <= 26 {
				res += dfs(i + 2)
			}
		}

		memo[i] = res
		return res
	}

	return dfs(0)
}

// Top dowm dp
// Time complexity: O(n)
// Space complexity: O(n)
