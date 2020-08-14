package main

/*
21. 合并两个有序链表
难度:简单
将两个升序链表合并为一个新的 升序 链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。
示例：
输入：1->2->4, 1->3->4
输出：1->1->2->3->4->4
*/
// ListNode ...
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{}
	tmpNode := head
	for l1 != nil && l2 != nil {
		// 如果l2大，那么保存l1，并且移动l1，再和l2对比
		if l1.Val < l2.Val {
			// 保存当前节点
			tmpNode.Next = l1
			l1 = l1.Next
		} else {
			// 保存当前节点
			tmpNode.Next = l2
			l2 = l2.Next
		}
		// 移动
		tmpNode = tmpNode.Next
	}

	// 任何一条循环完毕，保证另外一条合并进来
	if l1 != nil {
		tmpNode.Next = l1
	}
	if l2 != nil {
		tmpNode.Next = l2
	}
	// 丢弃初始化的第一个节点
	return head.Next
}
