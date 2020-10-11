package homework

// 917. 仅仅反转字母

func reverseOnlyLetters(s string) string {

	i := 0
	j := len(s) - 1
	// 将string转换为数组
	S := []rune(s)
	for {
		if i > j {
			break
		}

		if S[i] >= 'a' && S[i] <= 'z' || S[i] >= 'A' && S[i] <= 'Z' {
			if S[j] >= 'a' && S[j] <= 'z' || S[j] >= 'A' && S[j] <= 'Z' {
				S[i], S[j] = S[j], S[i]
				i++
				j--
			} else {
				j--
			}
		} else {
			i++
		}
	}
	return string(S)
}
