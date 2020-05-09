package stack

/*
844. 比较含退格的字符串
给定 S 和 T 两个字符串，当它们分别被输入到空白的文本编辑器后，判断二者是否相等，并返回结果。 # 代表退格字符。

注意：如果对空文本输入退格字符，文本继续为空。
 */

// 双指针遍历法
func BackspaceCompare(S string, T string) bool {
	sPtr := len(S)-1
	tPtr := len(T)-1
	sSkip := 0
	tSkip := 0
	for sPtr >=0 || tPtr >= 0 {
		for sPtr >= 0 {
			if S[sPtr] == '#' {
				sSkip++
				sPtr--
			} else if sSkip > 0 {
				sSkip--
				sPtr--
			} else {
				break
			}
		}
		for tPtr >= 0 {
			if T[tPtr] == '#' {
				tSkip++
				tPtr--
			} else if tSkip > 0 {
				tSkip--
				tPtr--
			} else {
				break
			}
		}
		if sPtr >= 0 && tPtr >= 0 && S[sPtr] != T[tPtr] {
			return false
		}
		if (sPtr >= 0) != (tPtr >= 0) {
			return false
		}
		sPtr--
		tPtr--
	}
	return true
}

// 栈方法
//func BackspaceCompare(S string, T string) bool {
//	s := backStack{}
//	t := backStack{}
//	for _, v := range S {
//		if v == '#' {
//			s.Pop()
//		} else {
//			s.Push(v)
//		}
//	}
//	for _, v := range T {
//		if v == '#' {
//			t.Pop()
//		} else {
//			t.Push(v)
//		}
//	}
//	for !s.IsEmpty() && !t.IsEmpty() {
//		if s.Pop() != t.Pop() {
//			return false
//		}
//	}
//	if !s.IsEmpty() || !t.IsEmpty() {
//		return false
//	}
//	return true
//}
//
//type backStack []int32
//
//func (s backStack) IsEmpty() bool {
//	return len(s) == 0
//}
//
//func (s *backStack) Push(val int32) {
//	*s = append(*s, val)
//}
//
//func (s *backStack) Pop() int32 {
//	old := *s
//	length := len(old)
//	if length == 0 {
//		return 0
//	}
//	*s = old[:length-1]
//	return old[length-1]
//}


