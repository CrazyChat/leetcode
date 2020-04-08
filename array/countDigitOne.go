package array

import (
	"math"
	"strconv"
)

/*
输入一个整数 n ，求1～n这n个整数的十进制表示中1出现的次数。
例如，输入12，1～12这些整数中包含1 的数字有1、10、11和12，1一共出现了5次。
示例 1：
输入：n = 12
输出：5

来源：力扣（LeetCode）43
链接：https://leetcode-cn.com/problems/1nzheng-shu-zhong-1chu-xian-de-ci-shu-lcof
 */

//func CountDigitOne(n int) int {
//	count, pow := 0, 1
//	for n >= pow {
//		count += n / (pow * 10) * pow
//		if incr := n%(pow*10) - pow + 1; incr > pow {
//			count += pow
//		} else if incr > 0 {
//			count += incr
//		}
//		pow *= 10
//	}
//	return count
//}

func CountDigitOne(n int) int {
	return f(n)
}

func f(n int) int {
	if n <= 0 {
		return 0
	}
	s := strconv.Itoa(n)
	high := int(s[0] - '0')
	pow := int(math.Pow10(len(s)-1))
	last := n - high * pow
	if high == 1 {
		return f(pow-1) + last + 1 + f(last)
	} else {
		return pow + high * f(pow-1) + f(last)
	}
}
