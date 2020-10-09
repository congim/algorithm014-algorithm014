package day5

// 371. 两整数之和
func getSum(a int, b int) int {
	if a == 0 {
		return b
	}
	if b == 0 {
		return a
	}

	var low, high int
	for {
		low = a ^ b
		high = a & b
		if high == 0 {
			break
		}
		a = low
		b = high << 1
	}
	return low
}
