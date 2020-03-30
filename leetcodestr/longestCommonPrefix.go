package leetcodestr

/*
编写一个函数来查找字符串数组中的最长公共前缀。
如果不存在公共前缀，返回空字符串 ""。
示例 1:
输入: ["flower","flow","flight"]
输出: "fl"

LeetCode 14
链接：https://leetcode-cn.com/problems/longest-common-prefix
 */
// 1 垂直扫描法
//func LongestCommonPrefix(strs []string) string {
//	length := len(strs)
//	if length == 0 { return "" }
//	prefix := strs[0]
//	for i := 0; i < length; i++ {
//		for strings.Index(strs[i], prefix) != 0 {
//			prefix = prefix[0:len(prefix)-1]
//			if len(prefix) == 0 {
//				return ""
//			}
//		}
//	}
//	return prefix
//}
// 2 水平扫描法
func LongestCommonPrefix(strs []string) string {
	length := len(strs)
	if length == 0 { return "" }
	for i := 0; i < len(strs[0]); i++ {
		c := strs[0][i]
		for j := 1; j < length; j++ {
			if len(strs[j]) == i || strs[j][i] != c {
				return strs[0][:i]
			}
		}
	}
	return strs[0]
}
