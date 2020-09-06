package homework

import "fmt"

/*
874. 模拟行走机器人
难度:简单
机器人在一个无限大小的网格上行走，从点 (0, 0) 处开始出发，面向北方。该机器人可以接收以下三种类型的命令：
-2：向左转 90 度
-1：向右转 90 度
1 <= x <= 9：向前移动 x 个单位长度
在网格上有一些格子被视为障碍物。

第 i 个障碍物位于网格点  (obstacles[i][0], obstacles[i][1])

机器人无法走到障碍物上，它将会停留在障碍物的前一个网格方块上，但仍然可以继续该路线的其余部分。

返回从原点到机器人所有经过的路径点（坐标为整数）的最大欧式距离的平方。
示例 1：

输入: commands = [4,-1,3], obstacles = []
输出: 25
解释: 机器人将会到达 (3, 4)
示例 2：

输入: commands = [4,-1,4,-2,4], obstacles = [[2,4]]
输出: 65
解释: 机器人在左转走到 (1, 8) 之前将被困在 (1, 4) 处
提示：

0 <= commands.length <= 10000
0 <= obstacles.length <= 10000
-30000 <= obstacle[i][0] <= 30000
-30000 <= obstacle[i][1] <= 30000
答案保证小于 2 ^ 31
*/

func robotSim(commands []int, obstacles [][]int) int {
	obstaclesMap := map[string]bool{}
	for _, v := range obstacles {
		str := fmt.Sprintf("%v %v", v[0], v[1])
		obstaclesMap[str] = true
	}

	maxFunc := func(n, m int) int {
		if n > m {
			return n
		}
		return m
	}

	dirs := [][]int{
		{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

	x, y, d, ans := 0, 0, 0, 0
	for _, v := range commands {
		if v == -1 { // right
			//  顺时针转
			d++
			if d == 4 {
				d = 0
			}
		} else if v == -2 { // left
			// 逆时针转
			d--
			if d == -1 {
				d = 3
			}
		} else {
			for i := 0; i < v; i++ {
				key := fmt.Sprintf("%v %v", x+dirs[d][0], y+dirs[d][1])
				if _, ok := obstaclesMap[key]; !ok {
					x += dirs[d][0]
					y += dirs[d][1]
				}
			}
		}
		ans = maxFunc(ans, x*x+y*y)
	}
	return ans
}
