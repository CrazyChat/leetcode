package meituan

import (
	"fmt"
	"sort"
)

/*
不能超过
时间限制：C/C++语言 1000MS；其他语言 3000MS
内存限制：C/C++语言 65536KB；其他语言 589824KB
题目描述：
给出一个序列包含n个正整数的序列A，你可以从中删除若干个数，使得剩下的数字中的最大值和最小值之差不超过x，请问最少删除多少个数字。

输入
输入第一行仅包含两个正整数n和x，表示给出的序列的长度和给定的正整数。(1<=n<=1000,1<=x<=10000)
接下来一行有n个正整数，即这个序列，中间用空格隔开。(1<=a_i<=10000)

输出
输出仅包含一个正整数，表示最少删除的数字的数量

样例输入
5 2
2 1 3 2 5
样例输出
1

提示
* 极端情况下，当删除到仅剩1个数时，最大值和最小值的差为0，故不会出现无解的情况。
 */

func MinDifference(nums []int, dif int) int {
	length := len(nums)
	// 特例判断
	if length <= 1 {
		fmt.Println(0)
		return 0
	}
	// 排序
	sort.Ints(nums)
	// 比较最大值和最小值的差值
	max := length - 1
	min := 0
	delCount := 0
	for min <= max && nums[max] - nums[min] > dif {
		if nums[min+1]  - nums[min] >  nums[max] - nums[max-1] {
			// 删除left
			min++
		} else {
				// 删除右边
				max--
		}
		delCount++
	}
	return delCount
}
