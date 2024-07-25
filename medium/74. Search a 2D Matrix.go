package medium

func searchMatrix(matrix [][]int, target int) bool {
	n := len(matrix)
	m := len(matrix[0])
	left, right := 0, n*m-1
	for left <= right {
		mid := (left + right) / 2
		r, c := mid/m, mid%m
		if matrix[r][c] == target {
			return true
		} else if matrix[r][c] < target {
			left++
		} else {
			right--
		}
	}

	return false
}

//Time: O(log(N*M))
//Space: O(1)
