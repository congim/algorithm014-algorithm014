package day7

// 我自己的题解
// 作者：cong-shao
// 链接：https://leetcode-cn.com/problems/container-with-most-water/solution/bao-li-mei-ju-fang-fa-by-cong-shao/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

func maxArea(height []int) int {
	// 枚举, 水的容量等于长*宽，然后记录下来max。看那一次最大
	var max int
	for i := 0; i < len(height); i++ {
		for j := i + 1; j < len(height); j++ {
			area := getArea(i, j, height)
			max = getMax(max, area)
		}
	}
	return max
}

func getArea(i, j int, height []int) int {
	if height[i] < height[j] {
		return height[i] * (j - i)
	}
	return height[j] * (j - i)
}

func getMax(old, new int) int {
	if new > old {
		return new
	}
	return old
}
