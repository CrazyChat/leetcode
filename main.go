package main

import (
	"fmt"
	"github.com/crazychat/leetcode/dynamic"
)

func main() {
	arr := []int{1,2,3,4}
	result := dynamic.MaxProfit(arr)
	fmt.Println(result)
}

func Bag(weight []int, value []int,  n, cap int) int {
	// 建立动态表格, // 初始化 -1
	state := make([][]int, n)
	for i := range state {
		state[i] = make([]int, cap+1)
	}
	for i := range state {
		for j := range state[i] {
			state[i][j] = -1
		}
	}
	fmt.Println(state)
	// 首行初始化
	state[0][0] = 0
	if weight[0] <= cap {
		state[0][weight[0]] = value[0]
	}
	// 遍历动态更改状态
	for i := 1; i < n; i++ {
		for j := 0; j <= cap; j++ {
			// 在上一层为true的基础上进行分支
			if state[i-1][j] != -1 {
				// 不放
				state[i][j] = state[i-1][j]
				// 放
				if j + weight[i] <= cap {
					state[i][j+weight[i]] = state[i-1][j] + value[i]
				}
			}
		}
	}
	// 输出最大值
	maxValue := -1
	for i := 0; i <= cap; i++ {
		if state[n-1][i] > maxValue {
			maxValue = state[n-1][i]
		}
	}
	return maxValue
}