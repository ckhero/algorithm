/**
 *@Description
 *@ClassName exists
 *@Date 2021/5/19 下午8:23
 *@Author ckhero
 */

package _2_exists

func exist(board [][]byte, word string) bool {
	cols := len(board)
	rows := len(board[0])
	for col := 0; col < cols; col ++ {
		for row := 0; row < rows; row++ {
			if DFS(board, col, row, word, 0) {
				return true
			}
		}
	}
	return false
}

func DFS(board [][]byte, col, row int, word string, index int) bool {
	cols := len(board)
	rows := len(board[0])
	if col < 0 || col >= cols || row < 0 || row >= rows || word[index] != board[col][row] {
		return false
	}

	if index == len(word) - 1 {
		return true
	}

	// 往四个方向找
	tmp := board[col][row]
	board[col][row] = '/'
	res := DFS(board, col - 1, row, word, index + 1 ) ||
		DFS(board, col + 1, row, word, index + 1 ) ||
		DFS(board, col, row - 1, word, index + 1 ) ||
		DFS(board, col, row + 1, word, index + 1 )
	board[col][row] = tmp
	return res
}
