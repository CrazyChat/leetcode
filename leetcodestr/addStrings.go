package leetcodestr

import (
	"fmt"
)

/*
给定两个字符串形式的非负整数 num1 和num2 ，计算它们的和。

注意：

num1 和num2 的长度都小于 5100.
num1 和num2 都只包含数字 0-9.
num1 和num2 都不包含任何前导零。
你不能使用任何內建 BigInteger 库， 也不能直接将输入的字符串转换为整数形式。

来源：力扣（LeetCode）415
链接：https://leetcode-cn.com/problems/add-strings
 */

func AddStrings(num1 string, num2 string) string {
	res := ""
	carry := 0
	temp := 0
	i := len(num1) - 1
	j := len(num2) - 1
	n1 := 0
	n2 := 0
	for i >= 0 || j >= 0 {
		if i >= 0 {
			n1 = int(num1[i]-'0')
		} else {
			n1 = 0
		}
		if j >= 0 {
			n2 = int(num2[j]-'0')
		} else {
			n2 = 0
		}
		temp = n1 + n2 + carry
		carry = temp/10
		res = fmt.Sprintf("%d%s", temp%10, res)
		i--
		j--
	}
	if carry == 1 {
		res = fmt.Sprintf("%d%s", 1, res)
	}
	return res
}
