package day2

// 387. 字符串中的第一个唯一字符
func firstUniqChar(s string) int {
	var lf [26]int
	for i, ch := range s {
		lf[ch-'a'] = i
	}
	for i, ch := range s {
		if i == lf[ch-'a'] {
			return i
		} else {
			lf[ch-'a'] = -1
		}
	}
	return -1
}
