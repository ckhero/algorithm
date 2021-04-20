/**
 *@Description
 *@ClassName build_tree
 *@Date 2021/4/9 上午9:14
 *@Author ckhero
 */

package main

import "fmt"

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func main()  {
	res := buildTree([]int{3,9,20,15,7}, []int{9,3,15,20,7})
	fmt.Println(res)
}

func buildTree(preorder []int, inorder []int) *TreeNode {

	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}

	currIndex := 0
	for k, val := range inorder {
		if val == preorder[0] {
			currIndex = k
			break
		}
	}

	return &TreeNode{
		Val:   preorder[0],
		Left:  buildTree(preorder[1:currIndex + 1], inorder[:currIndex]),
		Right: buildTree(preorder[currIndex + 1:], inorder[currIndex + 1:]),
	}
}

