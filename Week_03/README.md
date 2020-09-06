学习笔记

递归模板:

```
//终止条件
//处理当前层逻辑
//下个梦境
//处理当前层局部变量(非必须)
```

分治模板:
```
//终止条件
//处理当前层逻辑
//下个梦境
//合并最终结果
//处理当前层局部变量(非必须)
```
本周做的好几道题目,都是参考题解的思路,翻译过来的.唯一自己思考写出来的题目是 98.<<验证二叉搜索树>> , 附上代码
```
func isValidBST(root *TreeNode) bool {
	return helper_isValidBST(root, math.MinInt64, math.MaxInt64)
}

func helper_isValidBST(root *TreeNode, min, max int) bool {
	if root == nil {
		return true
	}
	//fmt.Println(root)
	if root.Val <= min || root.Val >= max {
		return false //min < val < max
	}
	l := helper_isValidBST(root.Left, min, root.Val)
	r := helper_isValidBST(root.Right, root.Val, max)
	return l && r //merge
}
```
自我学习总结: 分治法和回溯法是两种不同算法思路,分治顾名思义是"分而治之",大问题切分成小问题,逐个击破.回溯法是一种基于分治法的思路,在处理小问题之前,进行一次"裁剪"或者叫"预处理",减少递归次数的加强版分治.