package homework

// 205. 同构字符串

// 绝对匹配
func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	// 通过两个哈希表
	// 哈希表1：key：s对应元素，value：t对应元素
	// 哈希表2：key：t对应元素，value：s对应元素
	hm := make(map[rune]rune)
	hm_t := make(map[rune]rune)

	for k, v := range s {
		// 2.1 如果哈希表1已经存在，判断是否对应的t元素相等，否则返回失败
		if val, ok := hm[v]; ok {
			if rune(t[k]) != val {
				return false
			}
		} else {
			// 保存哈希表1 kv信息
			hm[v] = rune(t[k])
		}

		// 2.2 如果哈希表2对应元素已经存在，判断是否对应的s元素相等，否则返回失败
		if val, ok := hm_t[rune(t[k])]; ok {
			if val != v {
				return false
			}
		} else {
			// 保存哈希表2 kv信息
			hm_t[rune(t[k])] = v
		}
	}
	return true
}
