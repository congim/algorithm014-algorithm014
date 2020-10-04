package homework

import "sort"

// 56. 合并区间

func merge(intervals [][]int) [][]int {
	if intervals == nil {
		return [][]int{}
	}
	if len(intervals) == 0 {
		return [][]int{}
	}
	//先根据区间第一个值按小到大排序
	sort.Sort(intss(intervals))
	var result [][]int
	temp := intervals[0]
	i := 1
	for ; i < len(intervals); i++ {
		//输出temp ， 第二个变成temp
		if temp[1] < intervals[i][0] {
			result = append(result, temp)
			temp = intervals[i]
			continue
		}
		//temp吞掉第二个
		if temp[1] > intervals[i][0] && temp[1] > intervals[i][1] {
			continue
		}
		//temp合并第二个
		if temp[1] < intervals[i][1] {
			temp[1] = intervals[i][1]
			continue
		}
	}
	//最后一个特殊处理
	if i == len(intervals) {
		result = append(result, temp)
	}
	return result
}

type intss [][]int

func (s intss) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s intss) Less(i, j int) bool {
	return s[i][0] < s[j][0]
}

func (s intss) Len() int {
	return len(s)
}
