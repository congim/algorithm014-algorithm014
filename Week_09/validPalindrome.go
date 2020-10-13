package main

// 680. 验证回文字符串 Ⅱ
func validPalindrome(s string) bool {
	return isPalindrome(s, 0, len(s)-1, true)
}

func isPalindrome(s string, left, right int, count bool) bool {
	for ; left < right; left, right = left+1, right-1 {
		if s[left] != s[right] {
			return count && (isPalindrome(s, left+1, right, false) || isPalindrome(s, left, right-1, false))
		}
	}
	return true
}
