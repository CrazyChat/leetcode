package stack

import "math"

/*
155. 最小栈
设计一个支持 push ，pop ，top 操作，并能在常数时间内检索到最小元素的栈。

push(x) —— 将元素 x 推入栈中。
pop() —— 删除栈顶的元素。
top() —— 获取栈顶元素。
getMin() —— 检索栈中的最小元素。
 */

/*
法一：
type MinStack struct {
	Val []int
	Min []int
}
 */

// 法二
type MinStack struct {
	head *MinStackNode
}

type MinStackNode struct {
	val int
	min int
	next *MinStackNode
}

/** initialize your data structure here. */
// Constructor
func MinStackNodeConstructor() MinStack {
	return MinStack{&MinStackNode{min: math.MaxInt64, next: nil}}
}


func (this *MinStack) Push(x int)  {
	newNode := MinStackNode{ val: x, next: this.head }
	if x <= this.head.min {
		newNode.min = x
	} else {
		newNode.min = this.head.min
	}
	this.head = &newNode
}


func (this *MinStack) Pop()  {
	this.head = this.head.next
}


func (this *MinStack) Top() int {
	return this.head.val
}


func (this *MinStack) GetMin() int {
	return this.head.min
}
