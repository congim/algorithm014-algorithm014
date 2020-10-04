package day3

func reverseBits(num uint32) uint32 {
	var ans, zero uint32
	var bits uint32 = 31

	for num > 0 {
		r := num % 2
		ans = ans | (zero+r)<<bits
		bits--
		num = num / 2
	}
	return ans
}
