package structure

import "errors"

/*
 顺序栈
 */

type ArrStack struct {
	data []int
	//cap int	// 容量
}

func NewArrStack() ArrStack {
	s := ArrStack{[]int{}}
	return s
}

func (stack *ArrStack) Len() int {
	return len(stack.data)
}

func (stack *ArrStack) Push(value int) {
	stack.data = append(stack.data, value)
}

func (stack *ArrStack) Top() (int, error) {
	length := len(stack.data)
	if length == 0 {
		return 0, errors.New("the Stack is empty")
	}
	return stack.data[length-1], nil
}

func (stack *ArrStack) Pop() (int, error) {
	length := len(stack.data)
	if length == 0 {
		return 0, errors.New("the Stack is empty")
	}
	temp := stack.data[length-1]
	stack.data = stack.data[:length-1]
	return temp, nil
}

func (stack *ArrStack) IsEmpty() bool {
	if len(stack.data) == 0 {
		return true
	}
	return false
}

/*
 链式栈
 */
type LinkStack struct {
	top *StackNode
	len int
	//cap int	// 容量
}

type StackNode struct {
	value int
	next *StackNode
}

func NewLinkStack() LinkStack {
	return LinkStack{nil, 0}
}

func (stack *LinkStack) Len() int {
	return stack.len
}

func (stack *LinkStack) Push(value int) {
	node := StackNode{value, stack.top}
	stack.top = &node
	stack.len++
}

func (stack *LinkStack) Top() (int, error) {
	if stack.len == 0 {
		return -1, errors.New("the Stack is empty")
	}
	return stack.top.value, nil
}

func (stack *LinkStack) Pop() (int, error) {
	if stack.len == 0 {
		return -1, errors.New("the Stack is empty")
	}
	temp := stack.top.value
	stack.top = stack.top.next
	stack.len--
	return temp, nil
}

func (stack *LinkStack) IsEmpty() bool {
	if stack.len == 0 {
		return true
	}
	return false
}