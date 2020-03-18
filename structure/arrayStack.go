package structure

import "errors"

type ArrayStack struct {
	Items []stackValType
	Count int
}

func CreateArrayStack() *ArrayStack {
	new := ArrayStack{make([]stackValType, 0), 0}
	return &new
}

func (stack *ArrayStack) IsEmpty() bool {
	if stack.Count == 0 {
		return true
	}
	return false
}

func (stack *ArrayStack) Push(val stackValType) {
	stack.Items = append(stack.Items, val)
	stack.Count++
	return
}

func (stack *ArrayStack) Pop() (stackValType, error) {
	if stack.Count == 0 {
		return 0, errors.New("栈为空")
	}
	stack.Count--
	val := stack.Items[stack.Count]
	return val, nil
}

func (stack *ArrayStack) Top() (stackValType, error) {
	if stack.Count == 0 {
		return 0, errors.New("栈为空")
	}
	return stack.Items[stack.Count-1], nil
}
