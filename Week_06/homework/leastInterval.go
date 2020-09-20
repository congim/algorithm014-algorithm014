package homework

// 621. 任务调度器
func leastInterval(tasks []byte, n int) int {
	max := func(x, y int) int {
		if x < y {
			return y
		}
		return x
	}

	tmp, res := make(map[byte]int, 0), 0
	for _, v := range tasks {
		tmp[v]++
	}

	ans := 0
	for _, v := range tmp {
		ans = max(ans, v)
	}

	res = (ans - 1) * (n + 1)
	for _, v := range tmp {
		if v == ans {
			res++
		}
	}
	if len(tasks) <= res {
		return res
	}
	return len(tasks)
}
