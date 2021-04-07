/**
 *@Description
 *@ClassName level_order
 *@Date 2021/4/7 下午1:34
 *@Author ckhero
 */

package level_order_one

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	tmp := []*TreeNode{root}
	res := []int{}
	k := 1
	offset := 0

	for offset < k {
		for i := offset; i < k; i++ {
			currNode := tmp[i]
			if currNode.Left != nil {
				tmp = append(tmp, currNode.Left)
				k++
			}
			if currNode.Right != nil {
				tmp = append(tmp, currNode.Right)
				k++
			}
			res = append(res, currNode.Val)
			offset++
		}
	}
	return res
}
