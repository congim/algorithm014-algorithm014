package main

func intersect(nums1 []int, nums2 []int) []int {
	m1 := make(map[int]int)
	m2 := make(map[int]int)
	var res []int
	for _, value := range nums1 {
		m1[value]++
	}
	for _, value := range nums2 {
		m2[value]++
	}
	for key1, value1 := range m2 {
		if value2, ok := m1[key1]; ok {
			if value1 > value2 {
				for i := 0; i < value2; i++ {
					res = append(res, key1)
				}
			} else {
				for i := 0; i < value1; i++ {
					res = append(res, key1)
				}
			}
		}
	}
	return res
}

// 350. 两个数组的交集 II
// 难度:简单

// 给定两个数组，编写一个函数来计算它们的交集。
// 示例 1：
// 输入：nums1 = [1,2,2,1], nums2 = [2,2]
// 输出：[2,2]
// 示例 2:
// 输入：nums1 = [4,9,5], nums2 = [9,4,9,8,4]
// 输出：[4,9]

// 说明：
// 输出结果中每个元素出现的次数，应与元素在两个数组中出现次数的最小值一致。
// 我们可以不考虑输出结果的顺序。
// 进阶：
// 如果给定的数组已经排好序呢？你将如何优化你的算法？
// 如果 nums1 的大小比 nums2 小很多，哪种方法更优？
// 如果 nums2 的元素存储在磁盘上，内存是有限的，并且你不能一次加载所有的元素到内存中，你该怎么办？
// https://leetcode-cn.com/problems/intersection-of-two-arrays-ii/
