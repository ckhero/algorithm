/**
 *@Description
 *@ClassName level_order_two_test.go
 *@Date 2021/6/12 下午4:03
 *@Author ckhero
 */

package level_order_one

import (
	"fmt"
	"testing"
)

func TestLevelOrderTwo(t *testing.T) {
	node := &TreeNode{
		Val:   1,
		Left:  &TreeNode{
			Val:   2,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val:   3,
			Left:  nil,
			Right: nil,
		},
	}
	fmt.Println(levelOrderTwo(node))
}