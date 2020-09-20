package homework

// 221. 最大正方形
func maximalSquare(matrix [][]byte) int {
	if len(matrix) == 0 {
		return 0
	}
	m, n, res := len(matrix), len(matrix[0]), 0
	dp := make([][]int, m+1)
	for i := 0; i < m+1; i++ {
		dp[i] = make([]int, n+1)
	}
	for i := 1; i < m+1; i++ {
		for j := 1; j < n+1; j++ {
			if matrix[i-1][j-1] == '1' {
				dp[i][j] = min(dp[i-1][j], dp[i][j-1], dp[i-1][j-1]) + 1
				res = max(res, dp[i][j]*dp[i][j])
			}
		}
	}
	return res
}
func min(a, b, c int) int {
	if a > b {
		if b < c {
			return b
		}
		return c
	} else {
		if a < c {
			return a
		}
		return c
	}
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
