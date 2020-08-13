package main

import "log"

// 转换思路
// 1 -> 2 -> 3 -> 4 -> 5
// 2.Next = 1
// 1.Next = 2.Next
// loop (2.Next)
// return 2

// ListNode ...
type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	// 当前节点，前置节点
	current, prepositive := head, head.Next
	// 当前节点的next需要改变成它前置的next
	current.Next = swapPairs(prepositive.Next)
	// 前置的next要变成当前的节点
	prepositive.Next = current
	// 返回前置，因为要交换，前置已经变成当前节点了
	return prepositive
}

func main() {
	// 测试用例
	head := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val:  4,
					Next: nil,
				},
			},
		},
	}
	// 测试
	val := swapPairs(head)
	// 结果展示
	for val != nil {
		log.Println(val.Val)
		val = val.Next
	}
}
