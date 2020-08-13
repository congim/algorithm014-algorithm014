package main

/*
88. 合并两个有序数组
难度:简单
给你两个有序整数数组 nums1 和 nums2，请你将 nums2 合并到 nums1 中，使 nums1 成为一个有序数组。
说明:
初始化 nums1 和 nums2 的元素数量分别为 m 和 n 。
你可以假设 nums1 有足够的空间（空间大小大于或等于 m + n）来保存 nums2 中的元素。
示例:
输入:
nums1 = [1,2,3,0,0,0], m = 3
nums2 = [2,5,6],       n = 3
输出: [1,2,2,3,5,6]
*/

func merge(nums1 []int, m int, nums2 []int, n int) {
	// 反向操作，从尾部对比
	// 因为已经知道两个数组多长，且有序，nums1更能存放下
	// 那么从尾部对比查看哪个最大的应该放到最后
	// 最长下标 m + n - 1
	for i := (m + n - 1); i >= 0; i-- {
		if m > 0 && n > 0 {
			if nums1[m-1] > nums2[n-1] {
				// 存放最大的
				nums1[i] = nums1[m-1]
				m--
			} else {
				nums1[i] = nums2[n-1]
				n--
			}
		}
	}
	if m == 0 {
		for i := 0; i < n; i++ {
			nums1[i] = nums2[i]
		}
	}
}
