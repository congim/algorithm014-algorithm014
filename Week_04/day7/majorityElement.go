package day7

func majorityElement(nums []int) int {
	n := len(nums)

	if n == 0 {
		return -1
	}

	if n < 2 {
		return nums[0]
	}

	hash := make(map[int]int)

	for i := 0; i < n; i++ {
		hash[nums[i]]++
		count, _ := hash[nums[i]]
		if count > n/2 {
			return nums[i]
		}
	}

	return -1
}
