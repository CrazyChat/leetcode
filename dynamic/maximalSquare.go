package dynamic

import "fmt"

/*
221. 最大正方形
在一个由 0 和 1 组成的二维矩阵内，找到只包含 1 的最大正方形，并返回其面积。

示例:

输入:

1 0 1 0 0
1 0 1 1 1
1 1 1 1 1
1 0 0 1 0

输出: 4
 */

func MaximalSquare(matrix [][]byte) int {
	if len(matrix) == 0 {
		return 0
	}
	fmt.Println("matrix: ")
	for i := range matrix {
		fmt.Println(matrix[i])
	}
	h := len(matrix)
	w := len(matrix[0])
	dp := make([][]int, h)
	for i := range dp {
		dp[i] = make([]int, w)
	}
	var maxSquare int
	// 初始化边界
	for i := 0; i < w; i++ {
		if maxSquare != 1 && matrix[0][i] == '1' {
			maxSquare = 1
		}
		dp[0][i] = int(matrix[0][i]-'0')
	}
	for i := 0; i < h; i++ {
		if maxSquare != 1 && matrix[i][0] == '1' {
			maxSquare = 1
		}
		dp[i][0] = int(matrix[i][0]-'0')
	}
	fmt.Println("dp: ")
	for i := range dp {
		fmt.Println(dp[i])
	}
	/* 动态填表
	dp(i, j)=min(dp(i−1, j), dp(i−1, j−1), dp(i, j−1))+1
	 */
	for i := 1; i < h; i++ {
		for j := 1; j < w; j++ {
			if matrix[i][j] == '1' {
				dp[i][j] = MaximalSquareMin(dp[i-1][j], dp[i-1][j-1], dp[i][j-1]) + 1
				if dp[i][j] > maxSquare {
					maxSquare = dp[i][j]
				}
			}
		}
	}
	return maxSquare*maxSquare
}

func MaximalSquareMin(x, y, z int) int {
	min := x
	if y < min {
		min = y
	}
	if z < min {
		min = z
	}
	return min
}
