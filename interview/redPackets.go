package interview

import "fmt"

func RedPackets() {
	n, m, a, b := 0, 0, 0, 0
	_, _ = fmt.Scanf("%d %d %d %d", &n, &m, &a, &b)
	// 选择发红包还是买礼物花费少
	min := a
	if a > b {
		min = b
	}
	// 1. m == 0，直接买礼物或送红包
	if m == 0 {
		fmt.Println(n * min)
		return
	}
	// 2. n == m，不用花费
	if n == m {
		fmt.Println(0)
	}
	// 3. n > m, 只能发 n-m 个红包
	if n > m {
		fmt.Println((n - m) * a)
		return
	}
	// 4. n < m，需要考虑人去留
	fmt.Println(0)
}
