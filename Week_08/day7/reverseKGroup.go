package day7

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}
	dummyHead := &ListNode{}
	preNode := dummyHead
	stack := make([]*ListNode, 0)
	for head != nil {
		count := 0
		for head != nil && count < k {
			stack = append(stack, head)
			head = head.Next
			count++
		}
		l := len(stack)
		if l < k {
			preNode.Next = stack[0]
			break
		}
		for i := l - 1; i >= 0; i-- {
			preNode.Next = stack[i]
			preNode = stack[i]
		}
		preNode.Next = nil
		stack = stack[l:]
	}
	return dummyHead.Next
}
