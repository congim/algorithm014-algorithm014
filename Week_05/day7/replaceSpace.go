package day7

func replaceSpace(s string) string {
	var slice []rune
	for _, schar := range s {
		if schar == 32 {
			slice = append(slice, '%')
			slice = append(slice, '2')
			slice = append(slice, '0')
		} else {
			slice = append(slice, schar)
		}
	}

	return string(slice)
}
