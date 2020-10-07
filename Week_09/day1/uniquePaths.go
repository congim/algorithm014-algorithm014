package day1

// 62. 不同路径
func uniquePaths(m int, n int) int {
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
		for j := 0; j < n; j++ {
			if 0 == i || 0 == j {
				dp[i][j] = 1
			} else {
				dp[i][j] = dp[i][j-1] + dp[i-1][j]
			}
		}
	}
	return dp[m-1][n-1]
}
