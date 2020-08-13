package main

/*
26. 删除排序数组中的重复项
难度:简单
给定一个排序数组，你需要在 原地 删除重复出现的元素，使得每个元素只出现一次，返回移除后数组的新长度。
不要使用额外的数组空间，你必须在 原地 修改输入数组 并在使用 O(1) 额外空间的条件下完成。

示例 1:
给定数组 nums = [1,1,2],
函数应该返回新的长度 2, 并且原数组 nums 的前两个元素被修改为 1, 2。
你不需要考虑数组中超出新长度后面的元素。

示例 2:
给定 nums = [0,0,1,1,1,2,2,3,3,4],
函数应该返回新的长度 5, 并且原数组 nums 的前五个元素被修改为 0, 1, 2, 3, 4。
你不需要考虑数组中超出新长度后面的元素。
*/

// 双指针
func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	// 由于是有序数组，那么相同的元素必然是相邻的，只需要对比相邻或者相邻的下一个，然后loop下去即可
	slowIndex := 0
	fastIndex := 1
	for fastIndex < len(nums) {
		// 如果不相同，那么slowIndx+1赋值当前fastIndex的下标的值
		// 类似 n[0] != n[1]; 那么n[0+1] = n[1]
		// 0是可变动的
		if nums[slowIndex] != nums[fastIndex] {
			// 覆盖或者f[n] = f[n]的操作
			nums[slowIndex+1] = nums[fastIndex]
			// 前移
			slowIndex++
		}
		// 前移
		fastIndex++
	}
	// slowIndex从0开始，所以长度需要加1
	return slowIndex + 1
}
