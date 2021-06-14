/**
 *@Description
 *@ClassName inorder_traversal_test.go
 *@Date 2021/6/12 下午5:34
 *@Author ckhero
 */

package _4_inorder_traversal

import (
	"fmt"
	"testing"
)

func Test_inorderTraversal(t *testing.T) {
	node := &TreeNode{
		Val: 1,
		Right: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val:   3,
				Left:  nil,
				Right: nil,
			},
		},
	}
	fmt.Println(inorderTraversal(node))
}
