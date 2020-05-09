package main

import (
	"fmt"
	"math"
)

func main() {
	minStack := MinStack{}
	minStack.Push(-2)
	minStack.Push(0)
	minStack.Push(-3)
	fmt.Println(minStack.GetMin())
	minStack.Pop()
	fmt.Println(minStack.Top())
	fmt.Println(minStack.GetMin())
}

type MinStack struct {
	Val []int
	Min []int
}


/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{[]int{}, []int{}}
}


func (this *MinStack) Push(x int)  {
	this.Val = append(this.Val, x)
	if x <= this.GetMin() {
		this.Min = append(this.Min, x)
	}
}


func (this *MinStack) Pop()  {
	if this.Val[len(this.Val)-1] == this.GetMin() {
		this.Min = this.Min[:len(this.Min)-1]
	}
	this.Val = this.Val[:len(this.Val)-1]
}


func (this *MinStack) Top() int {
	return this.Val[len(this.Val)-1]
}


func (this *MinStack) GetMin() int {
	if len(this.Min) == 0 {
		return math.MaxInt64
	}
	return this.Min[len(this.Min)-1]
}




