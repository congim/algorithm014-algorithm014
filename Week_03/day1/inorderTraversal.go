package day1

// 94. 二叉树的中序遍历
// 左根右

// TreeNode ...
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var res []int
	rr(root, &res)
	return res
}

func rr(node *TreeNode, res *[]int) {
	if node == nil {
		return
	}
	rr(node.Left, res)
	*res = append(*res, node.Val)
	rr(node.Right, res)
}
