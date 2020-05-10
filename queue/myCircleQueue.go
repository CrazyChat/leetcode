package queue

/*
622. 循环队列
*/

type MyCircularQueue struct {
	head int
	tail int
	cap  int
	val  []int
}

/** Initialize your data structure here. Set the size of the queue to be k. */
func Constructor(k int) MyCircularQueue {
	return MyCircularQueue{cap: k + 1, val: make([]int, k+1)}
}

/** Insert an element into the circular queue. Return true if the operation is successful. */
func (this *MyCircularQueue) EnQueue(value int) bool {
	if (this.tail+1)%this.cap == this.head {
		return false
	}
	this.val[this.tail] = value
	this.tail = (this.tail + 1) % this.cap
	return true
}

/** Delete an element from the circular queue. Return true if the operation is successful. */
func (this *MyCircularQueue) DeQueue() bool {
	if this.head == this.tail {
		return false
	}
	this.head = (this.head + 1) % this.cap
	return true
}

/** Get the front item from the queue. */
func (this *MyCircularQueue) Front() int {
	if this.head == this.tail {
		return -1
	}
	return this.val[this.head]
}

/** Get the last item from the queue. */
func (this *MyCircularQueue) Rear() int {
	if this.head == this.tail {
		return -1
	}
	return this.val[(this.tail+this.cap-1)%this.cap]
}

/** Checks whether the circular queue is empty or not. */
func (this *MyCircularQueue) IsEmpty() bool {
	return this.head == this.tail
}

/** Checks whether the circular queue is full or not. */
func (this *MyCircularQueue) IsFull() bool {
	return (this.tail+1)%this.cap == this.head
}
