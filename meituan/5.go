package meituan

import (
	"fmt"
	"sort"
)

/*
max xor min
时间限制：C/C++语言 1000MS；其他语言 3000MS
内存限制：C/C++语言 65536KB；其他语言 589824KB
题目描述：
 给你一个长度为n的序列a，请你求出对每一个1<=l<r<=n的区间中最大值和最小值的异或和的异或和。

例如序列为{1,3,5},不同的a(1,2)=1^3,a(1,3)=1^5,a(2,3)=(3^5),a(1,2)^a(1,3)^a(2,3)=0，所以最后的答案是0。

输入
输入第一行仅包含一个正整数n，表示序列的长度。(1<=n<=10^5)

接下来一行有n个正整数a_i，表示序列a。(1<=a_i<=10^9)

输出
输出仅包含一个整数表示所求的答案。


样例输入
3
1 3 5
样例输出
0

规则
请尽量在全场考试结束10分钟前调试程序，否则由于密集排队提交，可能查询不到编译结果
点击“调试”亦可保存代码
编程题可以使用本地编译器，此页面不记录跳出次数
*/
func MaxxOrMin(arr []int) int {
	if len(arr) == 0 {
		return 0
	}
	sort.Ints(arr)
	// 奇数结果为0
	if len(arr) % 2 != 0 {
		return 0
	}
	// 偶数为所有数字的异或和
	result := arr[0]
	for i := 1; i < len(arr); i++ {
		fmt.Printf("%d ^ %d = ", result, arr[i])
		result = result ^ arr[i]
		fmt.Printf("%d\n", result)
	}
	return result
}