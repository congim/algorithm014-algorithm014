package day3

// 718. 最长重复子数组
func findLength(A []int, B []int) int {
	lenA := len(A)
	lenB := len(B)
	// 说明 B 比较短
	if lenB == min(lenA, lenB) {
		B = reverse(B)
		return helper(B, A)
	} else {
		A = reverse(A)
		return helper(A, B)
	}

}

// short, long
func helper(s, l []int) int {
	var res int
	// 滑窗核心
	i := 0
	for i = 0; i < len(l); i++ {
		tmpI := i
		tmpRes := 0
		// 溢出的话，就每次都从l 最后一位开始比较，但s就开始从第1，第2，，，开始比较了
		if i == len(l)-1 {
			tmpJ := 0

			for {
				if tmpJ >= len(s)-1 {
					break
				}
				tmpRes := 0
				tmpI = i
				for j := tmpJ; j <= i; j++ {
					if l[tmpI] == s[j] {
						tmpRes += 1
					} else {
						tmpRes = 0
					}
					tmpI -= 1
					if tmpRes > res {
						res = tmpRes
					}
				}
				tmpJ += 1
			}
			break
		}

		for j := 0; j <= i; j++ {
			if l[tmpI] == s[j] {
				tmpRes += 1
			} else {
				tmpRes = 0
			}
			tmpI -= 1
			if tmpRes > res {
				res = tmpRes
			}

		}
	}

	return res
}

func reverse(arr []int) []int {
	length := len(arr)
	for i := 0; i < length/2; i++ {
		temp := arr[length-1-i]
		arr[length-1-i] = arr[i]
		arr[i] = temp
	}
	return arr
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
