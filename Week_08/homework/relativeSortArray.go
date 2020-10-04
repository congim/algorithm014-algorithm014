package homework

import "sort"

// 1122. 数组的相对排序

func relativeSortArray(arr1 []int, arr2 []int) []int {
	m := make(map[int][]int)
	for _, v := range arr2 {
		m[v] = make([]int, 0)
	}
	var other []int
	for _, v := range arr1 {
		if l, ok := m[v]; ok {
			l = append(l, v)
			m[v] = l
		} else {
			other = append(other, v)
		}
	}
	sort.Ints(other)
	var res []int
	for _, v := range arr2 {
		res = append(res, m[v]...)
	}
	res = append(res, other...)
	return res
}
