package structure

import "fmt"

/*
155. 最小栈
设计一个支持 push，pop，top 操作，并能在常数时间内检索到最小元素的栈。

push(x) -- 将元素 x 推入栈中。
pop() -- 删除栈顶的元素。
top() -- 获取栈顶元素。
getMin() -- 检索栈中的最小元素。
示例:

MinStack minStack = new MinStack();
minStack.push(-2);
minStack.push(0);
minStack.push(-3);
minStack.getMin();   --> 返回 -3.
minStack.pop();
minStack.top();      --> 返回 0.
minStack.getMin();   --> 返回 -2.
 */

type MinStack struct {
	data []int
	minData []int
	dataLength int
	minDataLength int
}

func ConstructorMinStack() MinStack {
	return MinStack{make([]int, 0), make([]int, 0), 0, 0}
}
func (this *MinStack) Push(x int)  {
	if this.minDataLength == 0 || this.minData[this.minDataLength - 1] >= x {
		this.minData = append(this.minData, x)
		this.minDataLength++
	}
	this.data = append(this.data, x)
	this.dataLength++
	fmt.Println("插入", x, "后: ", this.minData)
}


func (this *MinStack) Pop()  {
	if this.dataLength > 0 {
		if this.data[this.dataLength-1] == this.minData[this.minDataLength-1] {	// 如果出栈的是最小值，minData也需要出栈
			this.minData = this.minData[0 : this.minDataLength-1]
			this.minDataLength--
		}
		this.data = this.data[0 : this.dataLength-1]
		this.dataLength--
	}
}

func (this *MinStack) Top() int {
	if this.dataLength > 0 {
		return this.data[this.dataLength-1]
	}
	return -1
}

func (this *MinStack) GetMin() int {
	if this.minDataLength > 0 {
		return this.minData[this.minDataLength-1]
	}
	return -1
}
