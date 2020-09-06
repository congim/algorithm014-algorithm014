package homework

import "math"

/*
433. 最小基因变化
难度:中等

一条基因序列由一个带有8个字符的字符串表示，其中每个字符都属于 "A", "C", "G", "T"中的任意一个。

假设我们要调查一个基因序列的变化。一次基因变化意味着这个基因序列中的一个字符发生了变化。

例如，基因序列由"AACCGGTT" 变化至 "AACCGGTA" 即发生了一次基因变化。

与此同时，每一次基因变化的结果，都需要是一个合法的基因串，即该结果属于一个基因库。

现在给定3个参数 — start, end, bank，分别代表起始基因序列，目标基因序列及基因库，请找出能够使起始基因序列变化为目标基因序列所需的最少变化次数。如果无法实现目标变化，请返回 -1。

注意:

起始基因序列默认是合法的，但是它并不一定会出现在基因库中。
所有的目标基因序列必须是合法的。
假定起始基因序列与目标基因序列是不一样的。
示例 1:

start: "AACCGGTT"
end:   "AACCGGTA"
bank: ["AACCGGTA"]

返回值: 1
示例 2:

start: "AACCGGTT"
end:   "AAACGGTA"
bank: ["AACCGGTA", "AACCGCTA", "AAACGGTA"]

返回值: 2
示例 3:

start: "AAAAACCC"
end:   "AACCCCCC"
bank: ["AAAACCCC", "AAACCCCC", "AACCCCCC"]

返回值: 3

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/minimum-genetic-mutation
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

var minCount int
var bankMap = map[string]bool{}

func minMutation(start string, end string, bank []string) int {
	if len(end) == 0 || len(start) == 0 || len(bank) == 0 {
		return -1
	}

	for _, v := range bank {
		bankMap[v] = true
	}
	if _, ok := bankMap[end]; !ok {
		return -1
	}

	minCount = math.MaxInt32

	mutationDfs(start, end, 0, map[string]bool{}, bank)
	if minCount == math.MaxInt32 {
		return -1
	}
	return minCount
}

func mutationDfs(start string, end string, count int, visited map[string]bool, bank []string) {
	if start == end {
		minCount = minTimes(minCount, count)
		return
	}

	length := len(bank[0])
	for i := 0; i < len(bank); i++ {
		diffCount := 0
		curMutation := []byte(bank[i])
		startMutation := []byte(start)
		for j := 0; j < length; j++ {
			if curMutation[j] != startMutation[j] {
				diffCount++
			}
			if diffCount > 1 {
				continue
			}
		}

		if diffCount == 1 {
			if _, ok := visited[bank[i]]; !ok {
				visited[bank[i]] = true
				mutationDfs(bank[i], end, count+1, visited, bank)
				delete(visited, bank[i])
			}
		}
	}
}

func minTimes(a, b int) int {
	if a >= b {
		return b
	}
	return a
}
