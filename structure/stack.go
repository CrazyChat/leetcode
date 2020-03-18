package structure

type stackValType int

//type Stack struct {
//	Val  stackValType
//	Next *Stack
//}

//func CreateStack() *Stack {
//	return &Stack{Next: nil}
//}
//
//func (stack *Stack) IsEmpty() bool {
//	if stack.Next == nil {
//		return true
//	}
//	return false
//}
//
//func (stack *Stack) Push(val stackValType) bool {
//	newStack := &Stack{ val, stack.Next}
//	stack.Next = newStack
//	return true
//}
//
//func (stack *Stack) Pop() (*Stack, error) {
//	if stack.Next == nil {
//		return nil, errors.New("stack is empty")
//	}
//	val := stack.Next
//	stack.Next = stack.Next.Next
//	return val, nil
//}
//
//func (stack *Stack) Top() (*Stack, error) {
//	if stack.Next == nil {
//		return nil, errors.New("stack is empty")
//	}
//	return stack.Next, nil
//}
