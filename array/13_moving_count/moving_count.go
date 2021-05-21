/**
 *@Description
 *@ClassName moving_count
 *@Date 2021/5/19 下午8:47
 *@Author ckhero
 */

package _3_moving_count

func movingCount(m int, n int, k int) int {
	var count int
	board := make([][]byte, m)
	for k, _ := range board {
		board[k] = make([]byte, n)
	}
	DFS(board, 0, 0, k, &count)
	return count
}

func DFS(board [][]byte, col, row, k int, count *int) {
	cols := len(board)
	rows := len(board[0])
	if sum(col, row) > k {
		return
	}
	if col < 0 || col >= cols || row < 0 || row >= rows || board[col][row] == '/' {
		return
	}
	board[col][row] = '/'
	*count++
	DFS(board, col-1, row, k, count)
	DFS(board, col+1, row, k, count)
	DFS(board, col, row-1, k, count)
	DFS(board, col, row+1, k, count)
}

func sum(m, n int) int {
	var res int
	for m != 0 {
		res += m % 10
		m /= 10
	}
	for n != 0 {
		res += n % 10
		n /= 10
	}
	return res
}
