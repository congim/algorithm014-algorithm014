package main

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for index, value := range nums {
		if rIndex, ok := m[target-value]; ok {
			return []int{index, rIndex}
		} else {
			m[value] = index
		}
	}
	return nil
}
