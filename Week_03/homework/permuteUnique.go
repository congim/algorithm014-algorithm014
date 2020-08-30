package homework

// 47. 全排列 II
func permuteUnique(nums []int) [][]int {
	res := [][]int{}
	dfs2(nums, &res, 0)
	return res
}

func dfs2(nums []int, res *[][]int, index int) {
	if index == len(nums) {
		*res = append(*res, dump(nums))
	}

	m := map[int]int{}
	for i := index; i < len(nums); i++ {
		if _, ok := m[nums[i]]; ok {
			continue
		}
		nums[i], nums[index] = nums[index], nums[i]
		dfs2(nums, res, index+1)
		nums[i], nums[index] = nums[index], nums[i]
		m[nums[i]] = 1
	}
}

func dump(a []int) []int {
	b := make([]int, len(a))
	copy(b, a)
	return b
}
