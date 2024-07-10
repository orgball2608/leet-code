package medium

func isValidSudoku(board [][]byte) bool {
	var rows [9][9]bool
	var columns [9][9]bool
	var squares [3][3][9]bool
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			cell := board[i][j]
			if cell == '.' {
				continue
			}

			// transform string number to number
			digit := int(cell-'0') - 1

			if rows[i][digit] || columns[j][digit] || squares[i/3][j/3][digit] {
				return false
			}

			rows[i][digit] = true
			columns[j][digit] = true
			squares[i/3][j/3][digit] = true
		}
	}
	return true
}

//Time: O(1)
//Space: O(1)
