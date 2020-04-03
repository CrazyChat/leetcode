package dynamic

/*
给定一个三角形，找出自顶向下的最小路径和。每一步只能移动到下一行中相邻的结点上。

例如，给定三角形：

[
     [2],
    [3,4],
   [6,5,7],
  [4,1,8,3]
]
自顶向下的最小路径和为 11（即，2 + 3 + 5 + 1 = 11）。

来源：力扣（LeetCode）120. 三角形最小路径和
链接：https://leetcode-cn.com/problems/triangle
 */
/*
从上往下
 */
//func MinimumTotal(triangle [][]int) int {
//	minLoadMap = map[[2]int]int{}
//	xLen := len(triangle)-1
//	return minLoad(triangle, 0, 0, xLen)
//}
////map 存储部分该点到底层最短路径
//var minLoadMap = make(map[[2]int]int)
////minLoad 计算该点出发到底部最小值
//func minLoad(triangle [][]int, x, y, xLen int) int {
//	if x == xLen {
//		return triangle[x][y]
//	}
//	leftMinLoad, ok1 := minLoadMap[[2]int{x+1, y}]
//	rightMinLoad, ok2 := minLoadMap[[2]int{x+1, y+1}]
//	if !ok1 {
//		leftMinLoad = minLoad(triangle, x+1, y, xLen)
//		// 存储
//		minLoadMap[[2]int{x+1, y}] = leftMinLoad
//	}
//	if !ok2 {
//		rightMinLoad = minLoad(triangle, x+1, y+1, xLen)
//		// 存储
//		minLoadMap[[2]int{x+1, y+1}] = rightMinLoad
//	}
//	return triangle[x][y] + min(leftMinLoad, rightMinLoad)
//}
//func min(x, y int) int {
//	if x < y {
//		return x
//	}
//	return y
//}

/*
从下往上
 */
func MinimumTotal(triangle [][]int) int {
	height := len(triangle)
	if height <= 0 {
		return 0
	}
	for i := height - 2; i >= 0; i--{
		for j := len(triangle[i])-1; j>=0; j--{
			triangle[i][j] = triangle[i][j] + min(triangle[i+1][j], triangle[i+1][j+1])
		}
	}
	return triangle[0][0]
}

func min(a,b int) int{
	if a>b{
		return b
	}

	return a
}