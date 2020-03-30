package leetcodestr

/*
给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。

示例 1:

输入: "abcabcbb"
输出: 3
解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。

来源：力扣（LeetCode）3
链接：https://leetcode-cn.com/problems/longest-substring-without-repeating-characters
 */
// 滑动队列
func lengthOfLongestSubstring(s string) int {
	length := len(s)
	if length == 0 {
		return 0
	}
	chardit := make(map[uint8]bool)
	i := 0
	final_count := 0
	count := 0
	for j := 0; j < length; j++ {
		if _, exists := chardit[s[j]]; exists {
			if count > final_count {
				final_count = count
			}
			for exists {
				count--
				delete(chardit, s[i])
				i++
				_, exists = chardit[s[j]]
			}
		}
		count++
		chardit[s[j]] = true
	}
	if count > final_count {
		final_count = count
	}
	return final_count
}