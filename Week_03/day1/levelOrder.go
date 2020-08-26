package day1

/*
429. N叉树的层序遍历
难度:中等
给定一个 N 叉树，返回其节点值的层序遍历。 (即从左到右，逐层遍历)。
例如，给定一个 3叉树 :
返回其层序遍历:
[
     [1],
     [3,2,4],
     [5,6]
]
*/

// Node ...
type Node struct {
	Val      int
	Children []*Node
}

var res [][]int

func levelOrder(root *Node) [][]int {
	if root == nil {
		return nil
	}
	res = [][]int{}
	dfs(root, 0)
	return res
}

func dfs(root *Node, level int) {
	if root == nil {
		return
	}
	if len(res) == level {
		res = append(res, []int{})
	}
	res[level] = append(res[level], root.Val)
	for _, n := range root.Children {
		dfs(n, level+1)
	}
}
