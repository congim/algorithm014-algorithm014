package day4

import "math"

// 633. 平方数之和
func judgeSquareSum(c int) bool {
	max := int(math.Sqrt(float64(c)))
	left, right := 0, max
	for left <= max {
		sum := left*left + right*right
		if sum == c {
			return true
		}
		if sum < c {
			left++
		} else {
			right--
		}
	}
	return false
}
