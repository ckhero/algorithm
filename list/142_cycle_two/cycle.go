/**
 *@Description
 *@ClassName cycle
 *@Date 2021/6/13 下午12:12
 *@Author ckhero
 */

package _42_cycle_two

type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	slow := head
	fast := head
	for fast != nil {
		slow = slow.Next
		if fast.Next == nil {
			return nil
		}
		fast = fast.Next.Next
		if slow != fast {
			continue
		}
		for slow != head {
			slow = slow.Next
			head = head.Next
		}
		return slow
	}
	return nil
}
