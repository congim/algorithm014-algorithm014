package day6

// 33. 搜索旋转排序数组
func search(nums []int, target int) int {
	low, high := 0, len(nums)-1
	for low < high-1 {
		mid := (high + low) / 2
		a, b, c := nums[low], nums[mid], nums[high]
		if b == target {
			return mid
		}
		switch {
		// 上面两个case在递增序列找同二分查找
		case target >= a && target <= b:
			high = mid
		case target >= b && target <= c:
			low = mid
		// 如果在递增的区间没有找到
		// 那么只可能在递减的区间中
		// 由于是旋转数组，只可能有一个递减区间
		case a > b:
			high = mid
		case b > c:
			low = mid
		default:
			return -1
		}
	}
	if nums[low] == target {
		return low
	}
	if nums[high] == target {
		return high
	}
	return -1
}
