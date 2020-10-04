package day5

// 198. 打家劫舍
func rob(nums []int) int {
	max := func(i, j int) int {
		if i > j {
			return i
		}
		return j
	}
	switch len(nums) {
	case 0:
		return 0
	case 1:
		return nums[0]
	case 2:
		return max(nums[0], nums[1])
	}
	pre, ans := 0, 0
	for _, i := range nums {
		ans, pre = max(ans, i+pre), ans
	}
	return ans
}
