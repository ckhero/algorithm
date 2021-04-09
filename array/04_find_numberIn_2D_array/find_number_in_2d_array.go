/**
 *@Description
 *@ClassName find_number_in_2d_array
 *@Date 2021/4/8 下午7:26
 *@Author ckhero
 */

package _4_find_numberIn_2D_array

func findNumberIn2DArray(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}
	row := 0
	col := len(matrix[0]) - 1
	for row < len(matrix) && col >= 0 {
		if matrix[row][col] == target {
			return true
		}

		if matrix[row][col] > target {
			col--
		} else {
			row++
		}
	}
	return false
}
