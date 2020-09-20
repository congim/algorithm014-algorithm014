package homework

// 647. 回文子串
func countSubstrings(s string) int {
	ans := len(s)
	length := len(s)
	var sym func(left, right int) int
	sym = func(left, right int) int {
		ans := 0
		for left >= 0 && right < length {
			if s[left] == s[right] {
				ans++
				left--
				right++
			} else {
				break
			}
		}
		return ans
	}

	for index := range s {
		ans += sym(index-1, index+1)
		ans += sym(index, index+1)
	}

	return ans
}
