package leetcodestr

import "fmt"

/*
给定两个字符串 s1 和 s2，写一个函数来判断 s2 是否包含 s1 的排列。

换句话说，第一个字符串的排列之一是第二个字符串的子串。

示例1:

输入: s1 = "ab" s2 = "eidbaooo"
输出: True
解释: s2 包含 s1 的排列之一 ("ba").

来源：力扣（LeetCode）567
链接：https://leetcode-cn.com/problems/permutation-in-string
 */

//滑动窗口
func CheckInclusion(s1, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}
	// 创建存储s1所有字符出现次数
	s1Arr := [26]int{}
	s1Len := len(s1)
	// 创建s2窗口字符出现次数
	s2Arr := [26]int{}
	s2Len := len(s2)
	for i := 0; i < s1Len; i++ {
		s1Arr[s1[i]-'a']++
		s2Arr[s2[i]-'a']++
	}
	fmt.Println("arr1: ", s1Arr)
	for left := 0; left < s2Len - s1Len; left++ {
		fmt.Println("arr2: ", s2Arr)
		if match(s1Arr, s2Arr) {
			return true
		}
		s2Arr[s2[left]-'a']--
		s2Arr[s2[left+s1Len]-'a']++
	}
	return match(s1Arr, s2Arr)
}

func match(arr1, arr2 [26]int) bool {
	for i := 0; i < 26; i++ {
		if arr1[i] != arr2[i] {
			return false
		}
	}
	return true
}
