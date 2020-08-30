package homework

// 77. 组合
func combine(n int, k int) [][]int {
	ret := [][]int{}
	if n <= 0 || k <= 0 || n < k {
		return ret
	}
	sl := make([]int, k)
	dfs(n, k, 1, sl, &ret)
	return ret
}

func dfs(n, k, start int, sl []int, ret *[][]int) {
	l := len(sl)
	for i := start; i <= n; i++ {
		if k == 1 {
			sl[l-1] = i
			dst := make([]int, l)
			copy(dst, sl)
			*ret = append(*ret, dst)
		} else {
			sl[l-k] = i
			dfs(n, k-1, i+1, sl, ret)
		}
	}
}
