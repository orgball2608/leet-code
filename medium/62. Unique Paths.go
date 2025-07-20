package medium

import "fmt"

func uniquePaths(m int, n int) int {
	memo := make(map[string]int)
	var dfs func(i, j int) int
	dfs = func(i, j int) int {
		if i == 0 || j == 0 {
			return 1
		}

		key := fmt.Sprintf("%d-%d", i, j)
		count, ok := memo[key]
		if ok {
			return count
		}

		temp := dfs(i-1, j) + dfs(i, j-1)
		memo[key] = temp
		return temp
	}

	return dfs(m-1, n-1)
}

// Top down approach using memoization
// Time Complexity: O(m * n)
// Space Complexity: O(m * n) for memoization

func uniquePaths2(m int, n int) int {
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}

	for i := 0; i < m; i++ {
		dp[i][0] = 1
	}
	for j := 0; j < n; j++ {
		dp[0][j] = 1
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}

	return dp[m-1][n-1]
}

// Bottom up approach using dynamic programming
// Time Complexity: O(m * n)
// Space Complexity: O(m * n) for the dp array

func uniquePaths3(m int, n int) int {
	dp := make([]int, n)
	for j := 0; j < n; j++ {
		dp[j] = 1
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[j] = dp[j] + dp[j-1]
		}
	}

	return dp[n-1]
}

// Optimized space complexity using a single array
// Time Complexity: O(m * n)
// Space Complexity: O(n) for the dp array
