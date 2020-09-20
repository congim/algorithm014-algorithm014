package homework

// 32. 最长有效括号
/**
1. 最长 关键字 可以联想到 使用 贪心 或者 动态规划
2. 思考此类问题如果不能一眼看出答案,可以先思考解空间搜索法:
    1)逐一遍历所有子串
    2)判断是否为有效括号
    3)存储最长值即可
3.  寻找重叠子问题发现,重复判断同一段子串是否有效,因此我们可以使用 dp数组缓存一段有效子串的长度
    那么分为几种情况:
        1) 当前字符为')'时,一定构不成有效子串直接忽略
        2) 当前字符为'('时
            a.s[i-1] 为'('
                则 i==0时 dp[i] = 2
                   i>0时  dp[i] = dp[i-1] + 2
            b. s[i-1] 为'('
                则  i - dp[i - 1] - 1 == 0 时
                      dp[i] = dp[i-1] + 2
                    i - dp[i - 1] - 1 != 0 时
                      dp[i] = dp[i-1] + 2 + dp[i-dp[i-1]-2]
4. 此题关键在于通过重叠问题确定dp数组的含义,并定义循环不变式,再次基础上 正确推导出 dp方程,然后转换为编码即可
*/

func longestValidParentheses(s string) int {
	dp, res := make([]int, len(s)), 0
	// dp 数组表示 当前位置i处 之前在s串中构造的有效括号长度
	for i, _ := range s {
		if i > 0 && s[i] == ')' {
			if s[i-1] == '(' {
				if i == 1 {
					dp[i] = 2
				} else {
					dp[i] = dp[i-2] + 2
				}
			} else if s[i-1] == ')' && i-dp[i-1]-1 >= 0 && s[i-dp[i-1]-1] == '(' {
				if i-dp[i-1]-1 == 0 {
					dp[i] = dp[i-1] + 2
				} else {
					dp[i] = dp[i-1] + 2 + dp[i-dp[i-1]-2]
				}
			}
			if dp[i] > res {
				res = dp[i]
			}
		}
	}
	return res
}
