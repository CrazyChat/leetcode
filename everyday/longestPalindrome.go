package everyday

/*
LeetCode 409
给定一个包含大写字母和小写字母的字符串，找到通过这些字母构造成的最长的回文串。
在构造过程中，请注意区分大小写。比如 "Aa" 不能当做一个回文字符串。
注意:
假设字符串的长度不会超过 1010。
示例 1:
输入:
"abccccdd"
输出:
7
解释:
我们可以构造的最长的回文串是"dccaccd", 它的长度是 7。
 */

/*
将问题转化为求解每个字符出现的次数，出现奇数的所有字符只能一个放在中间
所以，最长回文字符为：总长度-奇数字符+1
 */
func LongestPalindrome(s string) int {
	counts := make(map[uint8]int)
	length := len(s)
	if length == 0 {
		return 0
	} else if length == 1 {
		return 1
	}
	for i := 0; i < length; i++ {
		_, exists := counts[s[i]]
		if exists {
			counts[s[i]] += 1
		} else {
			counts[s[i]] = 1
		}
	}
	oddNum := 0
	for _, v := range counts {
		if v%2 != 0 {
			oddNum += 1
		}
	}
	if oddNum == 0 {
		return length
	}
	return length - oddNum + 1
}