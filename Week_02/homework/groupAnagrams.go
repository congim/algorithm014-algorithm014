package main

/*
49. 字母异位词分组
难度:中等
给定一个字符串数组，将字母异位词组合在一起。字母异位词指字母相同，但排列不同的字符串。
示例:
输入: ["eat", "tea", "tan", "ate", "nat", "bat"]
输出:
[
  ["ate","eat","tea"],
  ["nat","tan"],
  ["bat"]
]
说明：
所有输入均为小写字母。
不考虑答案输出的顺序。
*/

import "sort"

func groupAnagrams(strs []string) [][]string {
	result := make([][]string, 0)
	keyMap := make(map[string][]string)
	for _, str := range strs {
		key := getKey(str)
		keyMap[key] = append(keyMap[key], str)
	}
	for _, value := range keyMap {
		result = append(result, value)
	}
	return result
}

func getKey(str string) string {
	bytes := []byte(str)
	sort.Slice(bytes, func(i, j int) bool { return bytes[i] < bytes[j] })
	return string(bytes)
}
