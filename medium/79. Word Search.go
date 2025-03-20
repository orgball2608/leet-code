package medium

func exist(board [][]byte, word string) bool {
	var result bool

	var backtrack func(r, c, index int)
	backtrack = func(r, c, index int) {
		if index == len(word) {
			result = true
			return
		}

		if r < 0 || r > len(board)-1 || c < 0 || c > len(board[0])-1 || board[r][c] != word[index] {
			return
		}

		temp := board[r][c]
		board[r][c] = '$'
		backtrack(r, c+1, index+1)
		backtrack(r+1, c, index+1)
		backtrack(r, c-1, index+1)
		backtrack(r-1, c, index+1)
		board[r][c] = temp
	}

	for r := 0; r < len(board); r++ {
		for c := 0; c < len(board[0]); c++ {
			backtrack(r, c, 0)
			if result {
				return true
			}

		}
	}

	return false
}
