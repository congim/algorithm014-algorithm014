package homework

// 493. 翻转对

// Node ...
type Node struct {
	l, r       *Node
	val, count int
}

func reversePairs(nums []int) int {
	// BST：超时（数组如果单调递增，就退化为链表）
	ans, n := 0, len(nums)
	if n > 0 {
		root := &Node{nil, nil, nums[n-1], 1} // 1
		for i := n - 2; i >= 0; i-- {
			ans += queryBST(root, (nums[i]+1)>>1) // 2
			insertBST(root, nums[i])              // 3
		}
	}
	return ans
}
func insertBST(root *Node, val int) *Node {
	if root == nil { // 3.1
		return &Node{nil, nil, val, 1}
	}
	switch {
	case root.val == val: // 3.2
		root.count++
	case root.val < val: // 3.3
		root.r = insertBST(root.r, val)
	case root.val > val: // 3.4
		root.count++
		root.l = insertBST(root.l, val)
	}
	return root
}
func queryBST(root *Node, half int) int {
	if root == nil { // 2.1
		return 0
	}
	if root.val >= half { // 2.2
		return queryBST(root.l, half)
	} else { // 2.3
		return root.count + queryBST(root.r, half)
	}
}
