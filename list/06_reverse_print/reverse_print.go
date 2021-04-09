/**
 *@Description
 *@ClassName reverse_print
 *@Date 2021/4/8 下午7:39
 *@Author ckhero
 */

package _6_reverse_print

type ListNode struct {
	Val int
	Next *ListNode
}

func reversePrint(head *ListNode) []int {
	if head == nil {
		return []int{}
	}
	var dummyHead *ListNode
	for head != nil {
		next := head.Next
		head.Next = dummyHead
		dummyHead = head
		head = next
	}

	res := []int{}
	for dummyHead != nil {
		res = append(res, dummyHead.Val)
		dummyHead = dummyHead.Next
	}
	return res
}
