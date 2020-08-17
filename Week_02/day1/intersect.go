package day1

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
