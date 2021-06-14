/**
 *@Description
 *@ClassName inorder_traversal
 *@Date 2021/6/12 下午5:13
 *@Author ckhero
 */

package _4_inorder_traversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 迭代
func inorderTraversal(root *TreeNode) []int {
	res := []int{}
	stack := []*TreeNode{}
	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, root.Val)
		root = root.Right
	}
	return res
}

// 递归2222222
func inorderTraversalRecursion(root *TreeNode) []int {

	var res []int
	if root != nil {
		res = inorderTraversal(root.Left)
		res = append(res, root.Val)
		res = append(res, inorderTraversal(root.Right)...)
	}
	return res
}
