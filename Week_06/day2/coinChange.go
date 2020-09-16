package day2

import "sort"

/*322. 零钱兑换
难度:中等
给定不同面额的硬币 coins 和一个总金额 amount。编写一个函数来计算可以凑成总金额所需的最少的硬币个数。如果没有任何一种硬币组合能组成总金额，返回 -1。

示例 1:

输入: coins = [1, 2, 5], amount = 11
输出: 3
解释: 11 = 5 + 5 + 1
示例 2:

输入: coins = [2], amount = 3
输出: -1
*/
type Coins struct {
	min    int
	amount int
	coins  []int
}

func (c *Coins) findAmount(count int, cur int, index int) {
	if cur == 0 {
		if c.min > count || c.min < 0 {
			c.min = count
		}
		return
	}
	if index < 0 || cur < 0 {
		return
	}

	for i := cur / c.coins[index]; i >= 0; i-- {
		if i+count > c.min && c.min > 0 {
			break
		}
		c.findAmount(count+i, cur-i*c.coins[index], index-1)
	}
	return
}

func coinChange(coins []int, amount int) int {
	sort.Ints(coins)
	coin := &Coins{
		min:    -1,
		amount: amount,
		coins:  coins,
	}
	coin.findAmount(0, amount, len(coins)-1)
	return coin.min
}
