package easy

func generate(numRows int) [][]int {
	if numRows <= 0 {
		return [][]int{}
	}

	result := make([][]int, numRows)
	for i := 0; i < numRows; i++ {
		result[i] = make([]int, i+1)
		result[i][0] = 1
		result[i][i] = 1
	}

	for i := 2; i < numRows; i++ {
		for j := 1; j < i; j++ {
			result[i][j] = result[i-1][j] + result[i-1][j-1]
		}
	}

	return result
}

// 1
// 1 1
// 1 2 1
// 1 3 3 1
// 1 4 6 4 1

// Time: O(N^2)
// Space: O(N^2)
