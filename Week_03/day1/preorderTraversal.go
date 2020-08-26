package day1

// 144. 二叉树的前序遍历
// 根左右

func preorderTraversal(root *TreeNode) []int {
	var res []int
	ss(root, &res)
	return res
}

func ss(node *TreeNode, res *[]int) {
	if node != nil {
		*res = append(*res, node.Val)
		ss(node.Left, res)
		ss(node.Right, res)
	}
}
