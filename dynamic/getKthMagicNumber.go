package dynamic

/*
面试题 17.09. 第 k 个数
有些数的素因子只有 3，5，7，请设计一个算法找出第 k 个数。注意，不是必须有这些素因子，而是必须不包含其他的素因子。例如，前几个数按顺序应该是 1，3，5，7，9，15，21。
示例 1:
输入: k = 5
输出: 9
*/

func GetKthMagicNumber(k int) int {
	num := make([]int, k)
	var p3, p5, p7 int
	num[0] = 1
	for i := 1; i < k; i++ {
		num[i] = getKthMagicNumberMin(num[p3]*3, num[p5]*5, num[p7]*7)
		if num[i] == num[p3]*3 {
			p3++
		}
		if num[i] == num[p5]*5 {
			p5++
		}
		if num[i] == num[p7]*7 {
			p7++
		}
	}
	return num[k-1]
}

func getKthMagicNumberMin(a, b, c int) int {
	if b < a {
		a = b
	}
	if c < a {
		a = c
	}
	return a
}
