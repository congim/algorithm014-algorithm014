package day2

/*
560. 和为K的子数组
难度:中等
给定一个整数数组和一个整数 k，你需要找到该数组中和为 k 的连续的子数组的个数。

示例 1 :

输入:nums = [1,1,1], k = 2
输出: 2 , [1,1] 与 [1,1] 为两种不同的情况。
说明 :

数组的长度为 [1, 20,000]。
数组中元素的范围是 [-1000, 1000] ，且整数 k 的范围是 [-1e7, 1e7]。
通过次数69,396提交次数154,065
*/

func subarraySum(nums []int, k int) int {
	m := map[int]int{0: 1}
	var sum int
	var count int
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		if v, ok := m[sum-k]; ok {
			count += v
		}
		m[sum]++
	}
	return count
}
