package day1

// 191. 位1的个数

// 时间复杂度：O(1) 依赖1的位数 num为uint32所以最多32位
func hammingWeight(num uint32) int {
	var count int
	for ; num > 0; count++ {
		num = num & (num - 1)
	}
	return count
}
