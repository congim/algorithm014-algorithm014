package main

import "math"

type MinStack struct {
	stack    []int
	minStack []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		stack:    []int{},
		minStack: []int{math.MaxInt64},
	}
}

func (this *MinStack) Push(x int) {
	this.stack = append(this.stack, x)
	top := this.minStack[len(this.minStack)-1]
	this.minStack = append(this.minStack, min(x, top))
}

func (this *MinStack) Pop() {
	this.stack = this.stack[:len(this.stack)-1]
	this.minStack = this.minStack[:len(this.minStack)-1]
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
	return this.minStack[len(this.minStack)-1]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

// type Node struct {
// 	value int
// 	next  *Node
// 	pre   *Node
// }

// type MinStack struct {
// 	size int
// 	head *Node
// 	tail *Node
// }

// /** initialize your data structure here. */
// func Constructor() MinStack {
// 	head := &Node{}
// 	tail := &Node{}

// 	head.next = tail
// 	tail.pre = head

// 	return MinStack{
// 		size: 0,
// 		head: head,
// 		tail: tail,
// 	}
// }

// func (this *MinStack) Push(x int) {
// 	node := &Node{
// 		value: x,
// 	}

// 	node.pre = this.tail.pre
// 	this.tail.pre.next = node
// 	node.next = this.tail
// 	this.tail.pre = node
// 	this.size++
// }

// func (this *MinStack) Pop() {
// 	if this.size == 0 {
// 		return
// 	}
// 	this.tail.pre.next = this.tail
// 	this.tail.pre = this.tail.pre.pre

// 	this.size--
// }

// func (this *MinStack) Top() int {
// 	if this.size == 0 {
// 		return 0
// 	}
// 	return this.tail.pre.value
// }

// func (this *MinStack) GetMin() int {
// 	if this.size == 0 {
// 		return 0
// 	}
// 	tmp := this.head.next
// 	min := tmp.value
// 	for tmp.next != this.tail {
// 		if min > tmp.value {
// 			min = tmp.value
// 		}
// 	}
// 	return min
// }

// // MinStack ...
// type MinStack struct {
// 	size   int
// 	statck []int
// }

// /** initialize your data structure here. */

// // Constructor ...
// func Constructor() MinStack {
// 	return MinStack{
// 		size:   0,
// 		statck: make([]int, 0),
// 	}
// }

// func (this *MinStack) Push(x int) {
// 	this.size++
// 	this.statck = append(this.statck, x)
// }

// func (this *MinStack) Pop() {
// 	if this.size == 0 {
// 		return
// 	}
// 	this.statck = append(this.statck[:this.size-1], this.statck[this.size:]...)
// 	this.size--
// }

// func (this *MinStack) Top() int {
// 	if this.size == 0 {
// 		return 0
// 	}
// 	return this.statck[this.size-1]
// }

// func (this *MinStack) GetMin() int {
// 	var min int
// 	for index, value := range this.statck {
// 		if index == 0 {
// 			min = this.statck[0]
// 			continue
// 		}
// 		if min > value {
// 			min = value
// 		}
// 	}
// 	return min
// }

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */

/*

155. 最小栈

设计一个支持 push ，pop ，top 操作，并能在常数时间内检索到最小元素的栈。

push(x) —— 将元素 x 推入栈中。
pop() —— 删除栈顶的元素。
top() —— 获取栈顶元素。
getMin() —— 检索栈中的最小元素。


示例:
输入：
["MinStack","push","push","push","getMin","pop","top","getMin"]
[[],[-2],[0],[-3],[],[],[],[]]

输出：
[null,null,null,null,-3,null,0,-2]

解释：
MinStack minStack = new MinStack();
minStack.push(-2);
minStack.push(0);
minStack.push(-3);
minStack.getMin();   --> 返回 -3.
minStack.pop();
minStack.top();      --> 返回 0.
minStack.getMin();   --> 返回 -2.


提示：
pop、top 和 getMin 操作总是在 非空栈 上调用。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/min-stack
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
