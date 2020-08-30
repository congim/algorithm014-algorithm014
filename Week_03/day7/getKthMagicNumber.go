package day7

/*
面试题 17.09. 第 k 个数
难度:中等
有些数的素因子只有 3，5，7，请设计一个算法找出第 k 个数。注意，不是必须有这些素因子，而是必须不包含其他的素因子。例如，前几个数按顺序应该是 1，3，5，7，9，15，21。
示例 1:
输入: k = 5
输出: 9
*/
func getKthMagicNumber(k int) int {
	ans, a, b, c := []int{1}, 0, 0, 0
	for i := 1; i < k; i++ {
		ans = append(ans, min(ans[a]*3, ans[b]*5, ans[c]*7))
		if ans[i] == ans[a]*3 {
			a++
		}
		if ans[i] == ans[b]*5 {
			b++
		}
		if ans[i] == ans[c]*7 {
			c++
		}
	}
	return ans[k-1]
}

func min(x, y, z int) int {
	r := x
	if y < r {
		r = y
	}
	if z < r {
		r = z
	}
	return r
}
