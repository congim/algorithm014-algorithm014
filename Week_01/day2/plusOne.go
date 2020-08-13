/*
66. 加一
难度
简单

给定一个由整数组成的非空数组所表示的非负整数，在该数的基础上加一。

最高位数字存放在数组的首位， 数组中每个元素只存储单个数字。

你可以假设除了整数 0 之外，这个整数不会以零开头。

示例 1:

输入: [1,2,3]
输出: [1,2,4]
解释: 输入数组表示数字 123。
示例 2:

输入: [4,3,2,1]
输出: [4,3,2,2]
解释: 输入数组表示数字 4321。
*/

package main

import "log"

func plusOne(digits []int) []int {
	length := len(digits)
	// 空情况
	if length == 0 {
		return []int{1}
	}

	for i := length - 1; i >= 0; i-- {
		// 非9不用进位，直接返回即可
		if digits[i] != 9 {
			digits[i]++
			return digits
		}
		// 如果是9的情况加一该位就为0，让后进一位，又循环给i--位加一。直到i=0位。
		digits[i] = 0
	}
	// 如果0位加进一位那么要改变数组长度,并且首位需要变成1
	return append([]int{1}, digits...)
}

func main() {
	array := plusOne([]int{999})
	array2 := plusOne([]int{899})
	log.Println(array, array2)
}
