package homework

// 338. 比特位计数
func countBits(num int) []int {
	d := make([]int, num+1)
	d[0] = 0
	for i := 1; i <= num; i++ {
		if i&1 == 1 {
			d[i] = d[i-1] + 1
		} else {
			d[i] = d[i/2]
		}
	}
	return d
}
