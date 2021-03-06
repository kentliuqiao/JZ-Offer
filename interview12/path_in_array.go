package interview12

func hasPath(board [][]byte, word string) bool {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if dfs(board, i, j, word, 0) {
				return true
			}
		}
	}

	return false
}

func dfs(board [][]byte, i, j int, word string, k int) bool {
	if i >= len(board) || i < 0 || j >= len(board[0]) || j < 0 ||
		word[k] != board[i][j] {

		return false
	}
	if k == len(word)-1 {
		return true
	}
	board[i][j] = ' '
	res := dfs(board, i+1, j, word, k+1) ||
		dfs(board, i, j+1, word, k+1) ||
		dfs(board, i-1, j, word, k+1) ||
		dfs(board, i, j-1, word, k+1)
	board[i][j] = word[k]

	return res
}
