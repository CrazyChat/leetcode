package stack

/*
232. 用栈实现队列
使用栈实现队列的下列操作：

push(x) -- 将一个元素放入队列的尾部。
pop() -- 从队列首部移除元素。
peek() -- 返回队列首部的元素。
empty() -- 返回队列是否为空。

只能使用标准的栈操作 -- 也就是只有 push to top, peek/pop from top, size, 和 is empty 操作是合法的
 */


// 首先实现一个栈
type QueueStack []int

func (s QueueStack) IsEmpty() bool {
	return len(s) == 0
}

func (s QueueStack) Size() int {
	return len(s)
}

func (s QueueStack) Top() int {
	if len(s) == 0 {
		return -1
	}
	return s[len(s)-1]
}

func (s *QueueStack) Push(val int) {
	old := *s
	*s = append(old, val)
}

func (s *QueueStack) Pop() int {
	length := len(*s)
	if length == 0 {
		return 0
	}
	old := *s
	*s = old[:length-1]
	return old[length-1]
}

type MyQueue struct {
	s1 QueueStack
	s2 QueueStack
}


/** Initialize your data structure here. */
func MyQueueConstructor() MyQueue {
	return MyQueue{[]int{}, []int{}}
}


/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int)  {
	// 转移s2到s1，然后再push(x)
	if this.s1.IsEmpty() {
		for !this.s2.IsEmpty() {
			this.s1.Push(this.s2.Pop())
		}
	}
	this.s1 = append(this.s1, x)
}


/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	if this.s2.IsEmpty() {
		for !this.s1.IsEmpty() {
			this.s2.Push(this.s1.Pop())
		}
	}
	tmp := this.s2[len(this.s2)-1]
	this.s2 = this.s2[:len(this.s2)-1]
	return tmp
}


/** Get the front element. */
func (this *MyQueue) Peek() int {
	if this.s2.IsEmpty() {
		for !this.s1.IsEmpty() {
			this.s2.Push(this.s1.Pop())
		}
	}
	return this.s2[len(this.s2)-1]
}


/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	return this.s1.IsEmpty() && this.s2.IsEmpty()
}
