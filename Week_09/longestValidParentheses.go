package main

// 32. 最长有效括号

func longestValidParentheses(s string) int {
	//遇到(就将其压栈
	//遇到)，先弹栈，然后若是栈空，则说明没有与其相匹配的(，这样就将当前下标压栈；若栈非空，则用当前下标-当前栈顶元素下标，即为目前count数
	//注意，一开始要先压一个-1进栈，防止s的首元素是)
	count := 0
	parStack := []int{-1}
	for i, item := range s {
		if item == '(' {
			parStack = append(parStack, i)
		} else {
			parStack = parStack[:len(parStack)-1]
			if len(parStack) == 0 {
				parStack = append(parStack, i)
			} else {
				count = max(count, i-parStack[len(parStack)-1])
			}
		}
	}
	return count
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

//测试用例：
//"()(()"
