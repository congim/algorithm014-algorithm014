package day6

import "strconv"

func getHint(secret string, guess string) string {
	a, b := 0, 0
	mapS := make([]int, 10)
	mapG := make([]int, 10)
	for i := range secret {
		tmp, charGuess := secret[i], guess[i]
		if tmp == guess[i] {
			a++
		} else {
			mapS[tmp-'0']++
			mapG[charGuess-'0']++
		}
	}
	for i := 0; i < 10; i++ {
		//找到重叠的
		b += min(mapS[i], mapG[i])
	}
	return strconv.Itoa(a) + "A" + strconv.Itoa(b) + "B"
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
