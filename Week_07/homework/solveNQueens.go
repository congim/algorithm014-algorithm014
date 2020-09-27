package homework

// 51. N 皇后
import "strings"

var ans [][]string

var cols map[int]bool
var pie map[int]bool
var na map[int]bool

func solveNQueens(n int) [][]string {
	if n <= 0 {
		return ans
	}
	ans = [][]string{}

	cols = make(map[int]bool, n)
	pie = make(map[int]bool, n)
	na = make(map[int]bool, n)

	dfs(n, 0, []int{})
	return ans
}

func dfs(n int, row int, path []int) {
	if row >= n {
		ans = append(ans, generateBoard(path, n))
		return
	}

	for col := 0; col < n; col++ {
		if cols[col] || pie[row+col] || na[row-col] {
			continue
		}
		cols[col] = true
		pie[row+col] = true
		na[row-col] = true
		dfs(n, row+1, append(path, col))
		cols[col] = false
		pie[row+col] = false
		na[row-col] = false
	}
	return
}

func generateBoard(path []int, n int) []string {
	var res []string
	for _, col := range path {
		res = append(res, strings.Repeat(".", col)+"Q"+strings.Repeat(".", n-col-1))
	}
	return res
}
