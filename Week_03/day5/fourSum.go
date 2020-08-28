package day5

import "sort"

/*
18. 四数之和
难度:中等
给定一个包含 n 个整数的数组 nums 和一个目标值 target，判断 nums 中是否存在四个元素 a，b，c 和 d ，使得 a + b + c + d 的值与 target 相等？找出所有满足条件且不重复的四元组。
注意：
答案中不可以包含重复的四元组。
示例：
给定数组 nums = [1, 0, -1, 0, -2, 2]，和 target = 0。
满足要求的四元组集合为：
[
  [-1,  0, 0, 1],
  [-2, -1, 1, 2],
  [-2,  0, 0, 2]
]
*/

func fourSum(nums []int, target int) [][]int {
	ret := make([][]int, 0)
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	if len(nums) < 4 {
		return ret
	}

	for k := 0; k < len(nums)-3; k++ {
		for p := k + 1; p < len(nums)-2; p++ {
			for i, j := p+1, len(nums)-1; i < j; {
				s := nums[k] + nums[p] + nums[i] + nums[j]
				if s > target {
					j--
				} else if s < target {
					i++
				} else {
					ret = append(ret, []int{nums[k], nums[p], nums[i], nums[j]})
					for i < j && nums[i] == nums[i+1] {
						i++
					}
					for i < j && nums[j] == nums[j-1] {
						j--
					}
					i++
					j--
				}
			}
			for p < len(nums)-3 && nums[p] == nums[p+1] {
				p++
			}
		}
		for k < len(nums)-4 && nums[k] == nums[k+1] {
			k++
		}
	}
	return ret
}
