package array

/*
给定一个包含了一些 0 和 1 的非空二维数组 grid 。
一个 岛屿 是由一些相邻的 1 (代表土地) 构成的组合，这里的「相邻」要求两个 1 必须在水平或者竖直方向上相邻。你可以假设 grid 的四个边缘都被 0（代表水）包围着。
找到给定的二维数组中最大的岛屿面积。(如果没有岛屿，则返回面积为 0 。)
示例 1:
[[0,0,1,0,0,0,0,1,0,0,0,0,0],
 [0,0,0,0,0,0,0,1,1,1,0,0,0],
 [0,1,1,0,1,0,0,0,0,0,0,0,0],
 [0,1,0,0,1,1,0,0,1,0,1,0,0],
 [0,1,0,0,1,1,0,0,1,1,1,0,0],
 [0,0,0,0,0,0,0,0,0,0,1,0,0],
 [0,0,0,0,0,0,0,1,1,1,0,0,0],
 [0,0,0,0,0,0,0,1,1,0,0,0,0]]
对于上面这个给定矩阵应返回 6。注意答案不应该是 11 ，因为岛屿只能包含水平或垂直的四个方向的 1 。

来源：力扣（LeetCode）695
链接：https://leetcode-cn.com/problems/max-area-of-island
*/

func MaxAreaOfIsland(grid [][]int) int {
	res := 0
	yLength := len(grid)
	xLength := len(grid[0])
	gridTmp := make([][]int, len(grid))
	copy(gridTmp, grid)
	for i := 0; i < yLength; i++ {
		for j := 0; j < xLength; j++ {
			if grid[i][j] == 1 {
				// 以此为中心，向四周扩散
				res = max(res, dfs(i, j, gridTmp))
			}
		}
	}
	return res
}
func max(i, j int) int {
	if i >= j {
		return i
	}
	return j
}

func dfs(i, j int, grid [][]int) int {
	// 每次调用的时候默认num为1，进入后判断如果不是岛屿，则直接返回0，就可以避免预防错误的情况。
	if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[i]) || grid[i][j] == 0 {
		return 0
	}
	// 每次找到岛屿，则直接把找到的岛屿改成0，这是传说中的沉岛思想，就是遇到岛屿就把他和周围的全部沉默。
	grid[i][j] = 0
	num := 1
	num += dfs(i+1, j, grid)
	num += dfs(i-1, j, grid)
	num += dfs(i, j-1, grid)
	num += dfs(i, j+1, grid)
	return num
}

