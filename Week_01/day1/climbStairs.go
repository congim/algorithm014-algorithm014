package main

import "log"

/*
假设你正在爬楼梯。需要 n 阶你才能到达楼顶。

每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？

注意：给定 n 是一个正整数。

示例 1：

输入： 2
输出： 2
解释： 有两种方法可以爬到楼顶。
1.  1 阶 + 1 阶
2.  2 阶
示例 2：

输入： 3
输出： 3
解释： 有三种方法可以爬到楼顶。
1.  1 阶 + 1 阶 + 1 阶
2.  1 阶 + 2 阶
3.  2 阶 + 1 阶

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/climbing-stairs
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// func climbStairs(n int) int {
// 	// 递归+结果缓存
// 	cache := make(map[int]int)
// 	return climbstairs(n, cache)
// }

// func climbstairs(n int, cache map[int]int) int {
// 	if num, ok := cache[n]; ok {
// 		return num
// 	}

// 	if n == 1 {
// 		cache[n] = 1
// 		return 1
// 	}
// 	if n == 2 {
// 		cache[n] = 2
// 		return 2
// 	}
// 	num := climbstairs(n-1, cache) + climbstairs(n-2, cache)
// 	cache[n] = num
// 	return num
// }

// 大规模的问题由若干个子问题组成
// 上 1 阶台阶： 1(1)
// 上 2 阶台阶： 2(1 1 , 2)
// 上 3 阶台阶： 3 (f(1) + f(2))
// 上 n 阶台阶： f(n) = f(n -1) + f(n -2)

// 动态规划
func climbStairs(n int) int {
	if n == 1 {
		return 1
	}
	dp := make([]int, n+1)
	dp[1] = 1
	dp[2] = 2
	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}

func main() {
	n := climbStairs(100)
	log.Println(n)
}
