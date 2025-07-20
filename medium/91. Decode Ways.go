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

func numDecodings2(s string) int {
	n := len(s)
	if n == 0 || s[0] == '0' {
		return 0
	}

	dp := make([]int, len(s)+1)
	dp[0] = 1
	dp[1] = 1
	for i := 2; i <= n; i++ {
		oneDigit := s[i-1] - '0'
		twoDigit := (s[i-2]-'0')*10 + (s[i-1] - '0')
		if oneDigit > 0 && oneDigit < 10 {
			dp[i] += dp[i-1]
		}

		if twoDigit >= 10 && twoDigit <= 26 {
			dp[i] += dp[i-2]
		}
	}

	return dp[n]
}

// Bottom up dp
// Time complexity: O(n)
// Space complexity: O(N)

func numDecodings3(s string) int {
	n := len(s)
	if n == 0 || s[0] == '0' {
		return 0
	}

	prev1, prev2 := 1, 1
	for i := 2; i <= n; i++ {
		curr := 0
		oneDigit := s[i-1] - '0'
		twoDigit := (s[i-2]-'0')*10 + (s[i-1] - '0')
		if oneDigit > 0 && oneDigit < 10 {
			curr += prev1
		}

		if twoDigit >= 10 && twoDigit <= 26 {
			curr += prev2
		}
		prev2 = prev1
		prev1 = curr

	}

	return prev1
}

// Bottom up dp with space optimization
// Time complexity: O(n)
// Space complexity: O(1)
