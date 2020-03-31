package leetcodestr

import (
	"fmt"
	"strings"
)

/*
给定一个字符串，逐个翻转字符串中的每个单词。

 

示例 1：

输入: "the sky is blue"
输出: "blue is sky the"
示例 2：

输入: "  hello world!  "
输出: "world! hello"
解释: 输入字符串可以在前面或者后面包含多余的空格，但是反转后的字符不能包括。
示例 3：

来源：力扣（LeetCode）151
链接：https://leetcode-cn.com/problems/reverse-words-in-a-string
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */

// 时间：O(N)
// 空间：O(N)
func ReverseWords(s string) string {
	s = strings.TrimSpace(s)
	fmt.Println(s)
	srcBytes := []byte(s)
	if len(srcBytes) == 0 {
		return ""
	}
	var dstBytes []byte
	i := len(s)-1
	for i >= 0 {
		for i >= 0 && srcBytes[i] == ' ' {
			i--
		}
		j := i
		for i >= 0 && srcBytes[i] != ' ' {
			i--
		}
		fmt.Printf("单词:%s, ",  srcBytes[i+1:j+1])
		dstBytes = append(dstBytes, srcBytes[i+1:j+1]...)
		if i > 0 {
			dstBytes = append(dstBytes, ' ')
		}
		fmt.Printf("句子:%s\n", dstBytes)
	}
	return string(dstBytes)
}
// 时间：O(N)
// 空间：O(1)
//func ReverseWords(s string) string {
//	// 去除前后空格
//	s = strings.TrimSpace(s)
//	if len(s) == 0 {
//		return ""
//	}
//	// 反转字符串
//	return s
//}
