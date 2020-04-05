package dynamic

import "fmt"

/*
给定一个数组，它的第 i 个元素是一支给定股票第 i 天的价格。
设计一个算法来计算你所能获取的最大利润。你可以尽可能地完成更多的交易（多次买卖一支股票）。
注意：你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。

示例 1:
输入: [7,1,5,3,6,4]
输出: 7
解释: 在第 2 天（股票价格 = 1）的时候买入，在第 3 天（股票价格 = 5）的时候卖出, 这笔交易所能获得利润 = 5-1 = 4 。
     随后，在第 4 天（股票价格 = 3）的时候买入，在第 5 天（股票价格 = 6）的时候卖出, 这笔交易所能获得利润 = 6-3 = 3 。

来源：力扣（LeetCode）122. 买卖股票的最佳时机 II
链接：https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-ii
 */
// 暴力回溯
//var res = 0
//func MaxProfit2(prices []int) int {
//	length := len(prices)
//	if length < 2 {
//		return 0
//	}
//	res = 0	// 防止缓存
//	dfs(prices, 0, length, res, false)
//	return res
//}
//
//// dfs (day: 第几天，profit: 当前收益，status: 是否持有股票)
//func dfs(prices []int, day, length, profit int, status bool) {
//	// 结束了
//	if day == length {
//		res = max(res, profit)
//		return
//	}
//	// 不操作
//	dfs(prices, day+1, length, profit, status)
//	// 没有股票：买入
//	if !status {
//		dfs(prices, day+1, length, profit-prices[day], !status)
//	} else {	// 有股票，卖出
//		dfs(prices, day+1, length, profit+prices[day], !status)
//	}
//}

// 贪心解法
//func MaxProfit2(prices []int) int {
//	res := 0
//	length := len(prices)
//	diff := 0
//	for i := 0; i < length-1; i++ {
//		diff = prices[i+1] - prices[i]
//		if diff > 0 {
//			res += diff
//		}
//	}
//	return res
//}

// 动态规划解法
func MaxProfit2(prices []int) int {
	length := len(prices)
	if length < 2 {
		return 0
	}
	/* 创建数组
	行: 那一天（具有前缀性质，即考虑了之前天数的收益）能获得的最大利润
	列: 那一天是持有股票，还是持有现金。0: cash; 1: stock
	 */
	dp := make([][]int, length)
	for i := range dp {
		dp[i] = make([]int, 2)
	}
	// 第一列的含义：什么都不做的收益
	// 第二列的含义：卖出或者买入的收益
	// 初始化
	dp[0][0] = 0
	dp[0][1] = -prices[0]
	fmt.Println("初始化：")
	for i := range dp  {
		fmt.Println(dp[i])
	}
	// 循环填表
	for i := 1; i < length; i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i])
	}
	fmt.Println("填表完成：")
	for i := range dp  {
		fmt.Println(dp[i])
	}
	return dp[length-1][0]
}