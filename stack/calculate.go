package stack

import (
	"strconv"
)

/*
224. 基本计算器
实现一个基本的计算器来计算一个简单的字符串表达式的值。

字符串表达式可以包含左括号 ( ，右括号 )，加号 + ，减号 -，非负整数和空格
 */

func Calculate(s string) int {
	length := len(s)
	numStack := numStack{}
	sign := 1	// 1加号，-1减号
	result := 0
	tmp := 0
	for i := 0; i < length; i++ {
		//fmt.Println(result, tmp, sign, numStack)
		switch s[i] {
		case ' ':
		case '+':
			result += sign * tmp
			sign = 1
			tmp = 0
		case '-':
			result += sign * tmp
			sign = -1
			tmp = 0
		case '(':
			numStack.Push(result)
			numStack.Push(sign)
			sign = 1
			result = 0
			tmp = 0
		case ')':
			result += sign * tmp
			result *= numStack.Pop()
			result += numStack.Pop()
			tmp = 0
			sign = 1
		default:
			j := i
			for ; j < length; j++ {
				if s[j] < '0' || s[j] > '9' {
					break
				}
			}
			tmp, _ = strconv.Atoi(s[i:j])
			i = j - 1
		}
	}
	return result + sign * tmp
}

type numStack []int

func (s numStack) IsEmpty() bool {
	return len(s) == 0
}

func (s *numStack) Push(val int) {
	*s = append(*s, val)
}

func (s *numStack) Pop() int {
	old := *s
	*s = old[:len(old)-1]
	return old[len(old)-1]
}


//func Calculate(s string) int {
//	numStack := numStack{}
//	oprator := calStack{}
//	length := len(s)
//	for i := 0; i < length; i++ {
//		fmt.Println(numStack, oprator)
//		switch s[i] {
//		case ' ':
//		case '+', '-':
//			if top := oprator.Top(); top == "+" || top == "-" {
//				op := oprator.Pop()
//				if op == "+" {
//					numStack.Push(numStack.Pop()+numStack.Pop())
//				} else {
//					numStack.Push(-(numStack.Pop()-numStack.Pop()))
//				}
//			}
//			oprator.Push(s[i:i+1])
//		case '(':
//			oprator.Push(s[i:i+1])
//		case ')':
//			op := oprator.Pop()
//			for op != "(" {
//				if op == "+" {
//					numStack.Push(numStack.Pop()+numStack.Pop())
//				} else {
//					numStack.Push(-(numStack.Pop()-numStack.Pop()))
//				}
//				op = oprator.Pop()
//			}
//		default:
//			j := i
//			for ; j < length; j++ {
//				if s[j] < '0' || s[j] > '9' {
//					break
//				}
//			}
//			tmp, _ := strconv.Atoi(s[i:j])
//			numStack.Push(tmp)
//			i = j-1
//		}
//	}
//	for !oprator.IsEmpty() {
//		op := oprator.Pop()
//		if op == "+" {
//			numStack.Push(numStack.Pop()+numStack.Pop())
//		} else {
//			numStack.Push(-(numStack.Pop()-numStack.Pop()))
//		}
//	}
//	return int(numStack.Pop())
//}
//
//type numStack []int
//
//func (s *numStack) Push(val int) {
//	*s = append(*s, val)
//}
//
//func (s *numStack) Pop() int {
//	old := *s
//	*s = old[:len(old)-1]
//	return old[len(old)-1]
//}
//
//type calStack []string
//
//func (s calStack) IsEmpty() bool {
//	return len(s) == 0
//}
//
//func (s calStack) Top() string {
//	if len(s) == 0 {
//		return ""
//	}
//	return s[len(s)-1]
//}
//
//func (s *calStack) Push(val string) {
//	*s = append(*s, val)
//}
//
//func (s *calStack) Pop() string {
//	old := *s
//	*s = old[:len(old)-1]
//	return old[len(old)-1]
//}
