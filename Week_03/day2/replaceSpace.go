package day2

/*
剑指 Offer 05. 替换空格
难度:简单
请实现一个函数，把字符串 s 中的每个空格替换成"%20"。
示例 1：
输入：s = "We are happy."
输出："We%20are%20happy."
限制：
0 <= s 的长度 <= 10000
*/

func Export() {
	replaceSpace("")
}
func replaceSpace(s string) string {

	var slice []rune
	for _, schar := range s {
		if schar == 32 {
			slice = append(slice, '%')
			slice = append(slice, '2')
			slice = append(slice, '0')
		} else {
			slice = append(slice, schar)
		}
	}

	return string(slice)
}
