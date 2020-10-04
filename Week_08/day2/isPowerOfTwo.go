package day2

// 231. 2的幂

func isPowerOfTwo(n int) bool {
	var cnt = 0
	for n > 0 {
		if n&1 == 1 {
			if cnt == 1 {
				return false
			}
			cnt++
		}
		n >>= 1
	}
	return cnt == 1
}
