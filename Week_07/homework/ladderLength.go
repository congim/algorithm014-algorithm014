package homework

// 127. 单词接龙
func ladderLength(beginWord string, endWord string, wordList []string) int {
	wordList = append(wordList, beginWord)
	start, end, n := len(wordList)-1, 0, len(wordList)
	for end < n && wordList[end] != endWord { // 找到endWord
		end++
	}
	if end == n { // 未找到
		return 0
	}
	graphtic := buildGraphtic(wordList) // 构建邻接矩阵图
	return getShortPath(graphtic, start, end)
}
func buildGraphtic(wordList []string) [][]int { // 构建图
	n := len(wordList)
	grapth := make([][]int, n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if isConnect(wordList[i], wordList[j]) { // 当只有一个单词不同时
				grapth[i] = append(grapth[i], j) // grapth[i]保存同一状态的不同单词的位于wordList的位置
			}
		}
	}
	return grapth
}
func isConnect(s1, s2 string) bool { // s1与s2只有一位不同
	diff, n := 0, len(s1)
	for i := 0; i < n && diff <= 1; i++ {
		if s1[i] != s2[i] {
			diff++
		}
	}
	return diff == 1
}
func getShortPath(grapth [][]int, start, end int) int { // 广度遍历
	queue := []int{start}                // 保存每一层节点
	visited := make([]bool, len(grapth)) // 节点已经被访问
	visited[start] = true
	path := 1
	for len(queue) != 0 {
		size := len(queue)
		path++
		for ; size > 0; size-- {
			cur := queue[0]
			queue = queue[1:]
			for _, next := range grapth[cur] {
				if next == end {
					return path
				}
				if visited[next] {
					continue
				}
				visited[next] = true
				queue = append(queue, next)
			}
		}
	}
	return 0
}
