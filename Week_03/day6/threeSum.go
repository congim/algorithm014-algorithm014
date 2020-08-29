package day6

/*
15. 三数之和
难度:中等
给你一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？请你找出所有满足条件且不重复的三元组。
注意：答案中不可以包含重复的三元组。
示例：
给定数组 nums = [-1, 0, 1, 2, -1, -4]，

满足要求的三元组集合为：
[
  [-1, 0, 1],
  [-1, -1, 2]
]
*/

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	res := [][]int{}
	sort.Ints(nums)
	cache := make(map[string][]int)
	for i := 0; i < len(nums)-2; i++ {
		tmpRes := twoSum(nums[i+1:], -nums[i])
		for _, t := range tmpRes {
			if _, ok := cache[fmt.Sprintf("%d%d%d", nums[i], t[0], t[1])]; ok {
				continue
			} else {
				t = append(t, nums[i])
				cache[fmt.Sprintf("%d%d%d", nums[i], t[0], t[1])] = t
				res = append(res, t)
			}
		}
	}
	return res
}

func twoSum(nums []int, target int) [][]int {
	res := [][]int{}
	cache := make(map[int]struct{})
	for i := 0; i < len(nums); i++ {
		if _, ok := cache[target-nums[i]]; ok {
			res = append(res, []int{nums[i], target - nums[i]})
			delete(cache, target-nums[i])
		} else {
			cache[nums[i]] = struct{}{}
		}
	}
	return res
}
