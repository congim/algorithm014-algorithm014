package main

/*589. N叉树的前序遍历
难度:简单

给定一个 N 叉树，返回其节点值的前序遍历。
例如，给定一个 3叉树 :

返回其前序遍历: [1,3,5,6,2,4]。
说明: 递归法很简单，你可以使用迭代法完成此题吗?
*/

// Definition for a Node.

// Node ...
type Node struct {
	Val      int
	Children []*Node
}

func preorder(root *Node) []int {

	ret := make([]int, 0)
	dfs(&ret, root)
	return ret
}

func dfs(ret *[]int, node *Node) {
	if node == nil {
		return
	}
	*ret = append(*ret, node.Val)
	for _, childNode := range node.Children {
		dfs(ret, childNode)
	}
}
