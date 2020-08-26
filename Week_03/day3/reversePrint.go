package day3

/*
剑指 Offer 06. 从尾到头打印链表
难度:简单
输入一个链表的头节点，从尾到头反过来返回每个节点的值（用数组返回）。
示例 1：
输入：head = [1,3,2]
输出：[2,3,1]
限制：
0 <= 链表长度 <= 10000
*/

// ListNode ...
type ListNode struct {
	Val  int
	Next *ListNode
}

func reversePrint(head *ListNode) []int {
	if head == nil {
		return nil
	}
	head = reverseLink(nil, head)
	res := make([]int, 0)
	for head != nil {
		res = append(res, head.Val)
		head = head.Next
	}
	return res
}

func reverseLink(pre, node *ListNode) *ListNode {
	if node == nil {
		return pre
	}
	next := node.Next
	node.Next = pre
	return reverseLink(node, next)
}
