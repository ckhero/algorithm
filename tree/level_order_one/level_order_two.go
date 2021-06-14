/**
 *@Description
 *@ClassName lvel_order_two
 *@Date 2021/6/12 下午3:59
 *@Author ckhero
 */

package level_order_one


func levelOrderTwo(root *TreeNode) []int {
	var (
		res = []int{}
		queue = []*TreeNode{root}
		offset = 0
		left = 1
	)
	for left > 0 {
		node := queue[offset]
		if node.Left != nil {
			queue = append(queue, node.Left)
			left++
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
			left++
		}
		res = append(res, node.Val)
		offset++
		left--
	}

	return  res
}
