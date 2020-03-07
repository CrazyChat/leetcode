package array

/*
给定一个由整数组成的非空数组所表示的非负整数，在该数的基础上加一。
最高位数字存放在数组的首位， 数组中每个元素只存储单个数字。
你可以假设除了整数 0 之外，这个整数不会以零开头
例如: [1,2,3] => [1,2,4], [4,3,2,1] => [4,3,2,2], [9,9] => [1,0,0]
 */

func PlusOne(digits []int) []int {
	digitsSize := 0
	digitsSize = len(digits)
	if digitsSize == 1 {
		if digits[0] == 9 {
			return []int{1,0}
		}
		digits[0] += 1
		return digits
	}
	for i := 1; digitsSize-i >= 0; i++ {
		if digits[digitsSize-i] != 9 {
			for j := 1; j < i; j++ {
				digits[digitsSize-j] = 0
			}
			digits[digitsSize-i] += 1
			return digits
		}
	}
	// 全是9的情况
	returnSize := make([]int, len(digits)+1, cap(digits)+1)
	returnSize[0] = 1
	return returnSize
}
