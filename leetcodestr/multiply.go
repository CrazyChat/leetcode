package leetcodestr

import "fmt"

/*
给定两个以字符串形式表示的非负整数 num1 和 num2，返回 num1 和 num2 的乘积，它们的乘积也表示为字符串形式。

示例 1:

输入: num1 = "2", num2 = "3"
输出: "6"
示例 2:

输入: num1 = "123", num2 = "456"
输出: "56088"

来源：力扣（LeetCode）43
链接：https://leetcode-cn.com/problems/multiply-strings
 */

func Multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	num1Len := len(num1)
	num2Len := len(num2)
	// 保证num1比num2长或者一样长
	if num1Len < num2Len {
		num1, num2 = num2, num1
		num1Len, num2Len = num2Len, num1Len
	}
	res := ""
	for i := num2Len-1; i >= 0; i-- {
		numChar := int(num2[i]-'0')
		zeroNum := num2Len-1-i	// 补多少个0
		temp := StringMultiplyDigit(num1, numChar, zeroNum)		// 字符和num1字符串相乘
		res = AddStrings(res, temp)
	}
	return res
}

func StringMultiplyDigit(n1 string, n2 int, zeroNum int) string {
	fmt.Printf("%s * %d\n", n1, n2)
	n1Len := len(n1)
	res := ""
	carry := 0
	temp := 0
	for i := n1Len-1; i >= 0; i-- {
		temp = int(n1[i]-'0') * n2 + carry
		carry = temp / 10
		res = fmt.Sprintf("%d%s", temp%10, res)
	}
	if carry != 0 {
		res = fmt.Sprintf("%d%s", carry, res)
	}
	for zeroNum > 0 {
		res += "0"
		zeroNum--
	}
	return res
}

