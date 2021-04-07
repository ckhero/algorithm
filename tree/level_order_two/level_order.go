/**
 *@Description
 *@ClassName level_order
 *@Date 2021/4/7 下午1:45
 *@Author ckhero
 */

package level_order_two

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	queue := []*TreeNode{root}
	res := [][]int{}
	level := 0
	for len(queue) > 0 {
		tmpQueue := []*TreeNode{}
		res = append(res, []int{})
		for _, currNode := range queue {
			res[level] = append(res[level], currNode.Val)
			if currNode.Left != nil {
				tmpQueue = append(tmpQueue, currNode.Left)
			}
			if currNode.Right != nil {
				tmpQueue = append(tmpQueue, currNode.Right)
			}
		}
		queue = tmpQueue
		level++
	}

	return res
}
