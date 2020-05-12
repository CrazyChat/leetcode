package queue

import (
	"fmt"
)

/*
862. 和至少为 K 的最短子数组
*/

func ShortestSubarray(A []int, K int) int {
	sum := make([]int, len(A)+1)
	for i := 0; i < len(A); i++ {
		sum[i+1] = sum[i] + A[i]
	}
	fmt.Println(sum)
	return 0
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
