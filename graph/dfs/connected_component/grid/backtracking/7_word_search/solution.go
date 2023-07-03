package word_search

/*
	problem: https://leetcode.com/problems/word-search
*/

var (
	dr      = []int{-1, 0, 1, 0}
	dc      = []int{0, 1, 0, -1}
	rTotal  = 0
	cTotal  = 0
	visited [][]bool
	goRight = []int{0, 1}
	goDown  = []int{1, 0}
	goLeft  = []int{0, -1}
	goUp    = []int{-1, 0}
)

func exist(board [][]byte, word string) bool {
	buildVisited(board)
	rTotal = len(board)
	cTotal = len(board[0])

	for r := 0; r < rTotal; r++ {
		for c := 0; c < cTotal; c++ {
			if board[r][c] == word[0] {
				return dfs(board, word, r, c, 0)
			}
		}
	}

	return false
}

func dfs(board [][]byte, word string, row, col int, index int) bool {
	// check all character in word are finded in board
	if index >= len(word) {
		return true
	}

	if row < 0 || row >= rTotal || col < 0 || col >= cTotal {
		return false
	}

	if visited[row][col] || board[row][col] != word[index] {
		return false
	}

	visited[row][col] = true
	if dfs(board, word, row+goRight[0], col+goRight[1], index+1) ||
		dfs(board, word, row+goDown[0], col+goDown[1], index+1) ||
		dfs(board, word, row+goLeft[0], col+goLeft[1], index+1) ||
		dfs(board, word, row+goUp[0], col+goUp[1], index+1) {
		return true
	}

	visited[row][col] = false

	return false
}

func buildVisited(board [][]byte) {
	n := len(board)
	visited = make([][]bool, n, n)
	for i := 0; i < n; i++ {
		visited[i] = make([]bool, len(board[0]), len(board[0]))
	}
}
