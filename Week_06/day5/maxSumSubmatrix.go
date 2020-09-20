package day5

// 363. 矩形区域不超过 K 的最大数值和
import "math"

func maxSumSubmatrix(matrix [][]int, k int) int {
	rowNum, colNum := len(matrix), len(matrix[0])
	//因为行远大于行，所以按行从左到右求前缀和
	for row := 0; row < rowNum; row++ {
		for col := 1; col < colNum; col++ {
			matrix[row][col] += matrix[row][col-1]
		}
	}
	result := math.MinInt32
	// 先固定左右边界，移动上下边界，求和
	for left := 0; left < colNum; left++ {
		for right := left; right < colNum; right++ {
			for top := 0; top < rowNum; top++ {
				sum := 0
				for bottom := top; bottom < rowNum; bottom++ {
					if left == 0 {
						sum += matrix[bottom][right]
					} else {
						sum += matrix[bottom][right] - matrix[bottom][left-1]
					}
					if sum <= k && sum > result {
						result = sum
					}
				}
			}
		}
	}
	return result
}
