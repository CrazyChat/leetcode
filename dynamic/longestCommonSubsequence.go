package dynamic

import "fmt"

/*
给定两个字符串 text1 和 text2，返回这两个字符串的最长公共子序列。

一个字符串的 子序列 是指这样一个新的字符串：它是由原字符串在不改变字符的相对顺序的情况下删除某些字符（也可以不删除任何字符）后组成的新字符串。
例如，"ace" 是 "abcde" 的子序列，但 "aec" 不是 "abcde" 的子序列。两个字符串的「公共子序列」是这两个字符串所共同拥有的子序列。
若这两个字符串没有公共子序列，则返回 0。

示例 1:
输入：text1 = "abcde", text2 = "ace"
输出：3
解释：最长公共子序列是 "ace"，它的长度为 3。

来源：力扣（LeetCode）1143
链接：https://leetcode-cn.com/problems/longest-common-subsequence
 */

func LongestCommonSubsequence(text1 string, text2 string) int {
	text1Len := len(text1)
	text2Len := len(text2)
	rect := make([][]int, text1Len)
	for i := 0; i < text1Len; i++ {
		rect[i] = make([]int, text2Len)
	}
	// 初始化边界
	if text1[0] == text2[0] {
		rect[0][0] = 1
	}
	for i := 1; i < text2Len; i++ {
		if text1[0] == text2[i] {
			rect[0][i] = 1
		} else {
			rect[0][i]  = rect[0][i-1]
		}
	}
	for i := 1; i < text1Len; i++ {
		if text2[0] == text1[i] {
			rect[i][0] = 1
		} else {
			rect[i][0] = rect[i-1][0]
		}
	}
	for i := range rect {
		fmt.Println(rect[i])
	}
	// 开始填充整个表格
	for i := 1; i < text1Len; i++ {
		for j := 1; j < text2Len; j++ {
			if text1[i] == text2[j] {
				rect[i][j] = rect[i-1][j-1] + 1
			} else {
				rect[i][j] = max(rect[i-1][j], rect[i][j-1])
			}
		}
	}
	for i := range rect {
		fmt.Println(rect[i])
	}
	// 找出最大值
	return rect[text1Len-1][text2Len-1]
}
func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
