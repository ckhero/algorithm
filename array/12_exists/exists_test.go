/**
 *@Description
 *@ClassName exists_test
 *@Date 2021/5/19 下午8:37
 *@Author ckhero
 */

package _2_exists

import (
	"fmt"
	"testing"
)

func TestDFS(t *testing.T) {
	board := [][]byte{
		{'a'},
	}
	word := "ab"

	fmt.Println(exist(board, word))
}
