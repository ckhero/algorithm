/**
 *@Description
 *@ClassName level_order
 *@Date 2021/4/7 下午1:45
 *@Author ckhero
 */

package level_order_two

import "container/list"

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
			if level & 1 != 0 {
				res[level] = append([]int{currNode.Val}, res[level]...)
			} else {
				res[level] = append(res[level], currNode.Val)
			}
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

func levelOrder2(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	queue := list.New()
	queue.PushBack(root)
	res := [][]int{}
	level := 0
	for queue.Len() > 0 {
		res = append(res, []int{})
		l := queue.Len()
		for l > 0  {
			l--
			if level & 1 == 0 {
				currNode := queue.Remove(queue.Front()).(*TreeNode)
				res[level] = append(res[level], currNode.Val)
				if currNode.Left != nil {
					queue.PushBack(currNode.Left)
				}
				if currNode.Right != nil {
					queue.PushBack(currNode.Right)
				}
			} else {
				currNode := queue.Remove(queue.Back()).(*TreeNode)
				res[level] = append(res[level], currNode.Val)
				if currNode.Right != nil {
					queue.PushFront(currNode.Right)
				}
				if currNode.Left != nil {
					queue.PushFront(currNode.Left)
				}
			}
		}
		level++
	}

	return res
}
